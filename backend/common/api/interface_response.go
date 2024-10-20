package api

import "net/http"

type IResponse interface {
	Send(http.ResponseWriter) bool
	GetErrors() interface{}
	GetData() interface{}
	GetCode() int
	GetIsError() bool
}
