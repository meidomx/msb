package api

import "github.com/julienschmidt/httprouter"

type HttpMethod int

const (
	HttpMethodGet     HttpMethod = 1
	HttpMethodPost    HttpMethod = 2
	HttpMethodPut     HttpMethod = 3
	HttpMethodDelete  HttpMethod = 4
	HttpMethodOptions HttpMethod = 5
	HttpMethodHead    HttpMethod = 6
	HttpMethodPatch   HttpMethod = 7
)

type RequestFormat int

const (
	RequestFormatUrlEncoded RequestFormat = 1
	RequestFormatJson       RequestFormat = 2
	RequestFormatMultipart  RequestFormat = 3
)

type HttpApiHandler interface {
	Name() string
	HttpMethods() []HttpMethod
	UrlMapping() string
	Handler() httprouter.Handle
}
