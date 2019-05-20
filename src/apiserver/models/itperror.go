package models

import (
	"fmt"
	"net/http"
)

type ITPError struct {
	errMessage string
	statusCode int
}

func (e *ITPError) HasNoError() bool {
	return e.Status() == 0 && e.errMessage == ""
}

func (e *ITPError) Error() string {
	return e.errMessage
}

func (e *ITPError) Status() int {
	return e.statusCode
}

func (e *ITPError) Notfound(target string, err error) {
	e.errMessage = fmt.Sprintf("Target: %s was not found, with detail: %+v.", target, err)
	e.statusCode = http.StatusNotFound
}

func (e *ITPError) Conflict(target string, err error) {
	e.errMessage = fmt.Sprintf("Target: %s was conflict, with detail: %+v", target, err)
	e.statusCode = http.StatusConflict
}

func (e *ITPError) InternalError(err error) {
	e.errMessage = fmt.Sprintf("Internal error occurred: %+v", err)
	e.statusCode = http.StatusInternalServerError
}
