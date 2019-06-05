package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

type AuthController struct {
	BaseController
}

// @Title Get
// @Description Returns token info registered from Keycloak.
// @Success 200 {string} 	Successful returned token info from Keycloak.
// @Failure 400 Bad request.
// @Failure 401 Unauthorized.
// @Failure 403 The resouce specified was forbidden to access.
// @Failure 404 The resource specified was not found.
// @Failure 500 Internal error occurred at server side.
// @router / [get]
func (ac *AuthController) Get() {}

func (ac *AuthController) Prepare() {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, configURL)
	if err != nil {
		ac.CustomAbort(http.StatusInternalServerError, fmt.Sprintf("Failed to create provider: %+v", err))
	}
	oauth2Config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
	oidcConfig := &oidc.Config{
		ClientID: clientID,
	}
	verifier := provider.Verifier(oidcConfig)
	if ac.GetString("state") != state {
		ac.CustomAbort(http.StatusBadRequest, "state did not match")
	}
	oauth2Token, err := oauth2Config.Exchange(ctx, ac.GetString("code"))
	if err != nil {
		ac.CustomAbort(http.StatusInternalServerError, "Failed to exchange token: "+err.Error())
	}
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		ac.CustomAbort(http.StatusInternalServerError, "No id_token field in oauth2 token.")
	}
	idToken, err := verifier.Verify(ctx, rawIDToken)
	if err != nil {
		ac.CustomAbort(http.StatusInternalServerError, "Failed to verify ID Token:"+err.Error())
	}
	resp := struct {
		OAuth2Token   *oauth2.Token
		IDTokenClaims *json.RawMessage
	}{oauth2Token, new(json.RawMessage)}
	if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
		ac.CustomAbort(http.StatusInternalServerError, err.Error())
	}
	ac.Data["json"] = resp
	ac.ServeJSON()
}
