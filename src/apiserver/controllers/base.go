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
			return
		}
		bc.CustomAbort(http.StatusInternalServerError, fmt.Sprintf("Error occurred: %+v", err))
	}
}
