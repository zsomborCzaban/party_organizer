package api

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	IsError bool        `json:"is_error"`
	Code    int         `json:"code"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

func (r *Response) GetErrors() interface{} {
	return r.Errors
}

func (r *Response) GetData() interface{} {
	return r.Data
}

func (r *Response) GetCode() int {
	return r.Code
}

func (r *Response) GetIsError() bool {
	return r.IsError
}

func (r *Response) Send(w http.ResponseWriter) bool {
	data, err := json.Marshal(r)
	if err != nil {
		return false
	}

	w.WriteHeader(r.Code)
	w.Header().Set("Content-Type", "application/json")
	_, e := w.Write(data)

	if e != nil {
		return false
	}

	return true
}

func Error(kind int, errors interface{}) IResponse {
	return &Response{
		IsError: true,
		Code:    kind,
		Errors:  errors,
	}
}

func Success(data interface{}) IResponse {
	return &Response{
		IsError: false,
		Code:    200,
		Data:    data,
	}
}

func ErrorValidation(errors interface{}) IResponse { return Error(http.StatusNotAcceptable, errors) }

func ErrorInternalServerError(errors interface{}) IResponse {
	return Error(http.StatusInternalServerError, errors)
}

func ErrorBadRequest(error string) IResponse {
	return Error(http.StatusBadRequest, error)
}

func ErrorNotFound(entity string) IResponse {
	return Error(http.StatusNotFound, entity+", was not found")
}

func ErrorInvalidCredentials() IResponse {
	ve := NewValidationErrors()
	ve.CollectValidationError("", "invalid password or username", nil)

	return Error(http.StatusNotAcceptable, ve.Errors)
}

func ErrorUnauthorized(message string) IResponse {
	return Error(http.StatusUnauthorized, message)
}
