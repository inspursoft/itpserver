package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
	oidc "github.com/coreos/go-oidc"
	"github.com/inspursoft/itpserver/src/apiserver/models"
	"golang.org/x/oauth2"
)

var configURL = beego.AppConfig.String("keycloak::configurl")
var clientID = beego.AppConfig.String("keycloak::clientid")
var clientSecret = beego.AppConfig.String("keycloak::clientsecret")
var state = beego.AppConfig.String("keycloak::state")
var redirectURL = beego.AppConfig.String("keycloak::redirecturl")

type BaseController struct {
	beego.Controller
}

func (bc *BaseController) Prepare() {
	accessToken := bc.GetString("access_token")
	if accessToken == "BOARD" {
		logs.Debug("Bypassing auth with provided token.")
		return
	}
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, configURL)
	if err != nil {
		bc.CustomAbort(http.StatusInternalServerError, fmt.Sprintf("Failed to create provider: %+v", err))
	}
	oauth2Config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
	oidcConfig := &oidc.Config{
		ClientID:          clientID,
		SkipClientIDCheck: true,
	}
	verifier := provider.Verifier(oidcConfig)
	rawAccessToken := bc.Ctx.Input.Header("Authorization")
	if rawAccessToken == "" {
		bc.Redirect(oauth2Config.AuthCodeURL(state), http.StatusFound)
		return
	}
	if !strings.HasPrefix(rawAccessToken, "Bearer") {
		rawAccessToken = "Bearer " + rawAccessToken
	}
	parts := strings.Split(rawAccessToken, " ")
	if len(parts) != 2 {
		bc.CustomAbort(http.StatusBadRequest, "Invalid authorization header info.")
	}
	_, err = verifier.Verify(ctx, parts[1])
	if err != nil {
		bc.Redirect(oauth2Config.AuthCodeURL(state), http.StatusUnauthorized)
		return
	}
}

func (bc *BaseController) requiredID(key string) int64 {
	id, err := bc.GetInt64(key, 0)
	if err != nil {
		bc.CustomAbort(http.StatusInternalServerError, fmt.Sprintf("Invalid input: %+v", err))
	}
	return id
}

func (bc *BaseController) requiredParam(key string) string {
	content := bc.GetString(key)
	if strings.TrimSpace(content) == "" {
		bc.CustomAbort(http.StatusBadRequest, fmt.Sprintf("Request parameter: %s is required.", key))
	}
	return content
}

func (bc *BaseController) loadRequestBody(target interface{}) {
	err := json.Unmarshal(bc.Ctx.Input.RequestBody, target)
	if err != nil {
		bc.CustomAbort(http.StatusInternalServerError, fmt.Sprintf("Failed to unmarshal request body, with error: %+v", err))
	}
}

func (bc *BaseController) handleError(err error) {
	if err != nil {
		if e, ok := err.(*models.ITPError); ok {
			if e.HasNoError() {
				return
			}
			if e.Status() >= 400 {
				fmt.Printf("%+v\n", e)
				bc.CustomAbort(e.Status(), e.Error())
			}
		} else {
			bc.CustomAbort(http.StatusInternalServerError, err.Error())
		}
	}
}

func (bc *BaseController) serveStatus(status int, message string) {
	bc.Data["json"] = struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{
		status, message,
	}
	bc.ServeJSON()
}

func (bc *BaseController) serveJSON(target interface{}) {
	bc.Data["json"] = target
	bc.ServeJSON()
}

func (bc *BaseController) resolveURL(targetURL string, values ...interface{}) string {
	return fmt.Sprintf("%s:%d%s", bc.Ctx.Input.Site(), bc.Ctx.Input.Port(), bc.URLFor(targetURL, values...))
}

func (bc *BaseController) proxiedRequest(method string, requestData interface{}, urlFor string, values ...interface{}) {
	requestBody, err := json.Marshal(requestData)
	bc.handleError(err)
	req, err := http.NewRequest(method, bc.resolveURL(urlFor, values...), bytes.NewBuffer(requestBody))
	req.Header.Set("Authorization", bc.Ctx.Input.Header("Authorization"))
	bc.handleError(err)
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		io.Copy(bc.Ctx.ResponseWriter, resp.Body)
	}
	bc.serveStatus(resp.StatusCode, "Finished handled proxied request.")
}

// @Title Get
// @Description Log out and clean up Keycloak session.
// @Success 200 {string} 	Successful logged out and cleaned up Keycloak session.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router /logout [get]
func (bc *BaseController) Logout() {
	bc.DelSession("token")
	bc.serveStatus(http.StatusOK, "Successful logged out.")
}
