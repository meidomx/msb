package api

import (
	"reflect"

	"github.com/julienschmidt/httprouter"
)

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
type ResponseFormat int

const (
	RequestFormatUrlEncoded RequestFormat = 1
	RequestFormatJson       RequestFormat = 2
	RequestFormatMultipart  RequestFormat = 3
)

const (
	ResponseFormatRawBinary ResponseFormat = 1
	ResponseFormatJson      ResponseFormat = 2
)

type HttpRequest struct {
	RequestURI string
	Params     httprouter.Params

	RequestObject interface{}
}

type HttpResponse struct {
	HttpStatus      int
	HttpContentType string

	HandleResult
}

type HttpApiSimpleHandler interface {
	Name() string
	HttpMethods() []HttpMethod
	ContentTypes() (RequestFormat, ResponseFormat)
	UrlMapping() string
	RequestType() reflect.Type
	Handle(request *HttpRequest) *HttpResponse
}
