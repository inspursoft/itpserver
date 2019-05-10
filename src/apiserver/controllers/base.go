package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/astaxie/beego"
	"github.com/inspursoft/itpserver/src/apiserver/models"
)

type BaseController struct {
	beego.Controller
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
			bc.CustomAbort(e.Status(), e.Error())
		}
		bc.CustomAbort(http.StatusInternalServerError, fmt.Sprintf("Error occurred: %+v", err))
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
