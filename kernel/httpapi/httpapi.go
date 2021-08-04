package httpapi

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime"
	"net"
	"net/http"
	"reflect"
	"strings"

	"github.com/meidomx/msb/api"
	"github.com/meidomx/msb/config"
	"github.com/meidomx/msb/utils"

	"github.com/julienschmidt/httprouter"
)

var _ContentTypeJson = mime.TypeByExtension("json")

type HttpApi struct {
	mux *httprouter.Router

	addr     string
	server   *http.Server
	listener net.Listener

	tmpMux *httprouter.Router
}

func NewHttpApi(cfg config.MsbConfig) *HttpApi {
	httpApi := new(HttpApi)
	httpApi.addr = cfg.Global.HttpApiAddr

	return httpApi
}

func (this *HttpApi) BuildHttpApiHandlers(hs []api.HttpApiSimpleHandler) {
	this.tmpMux = httprouter.New()
	for _, v := range hs {
		for _, m := range v.HttpMethods() {
			switch m {
			case api.HttpMethodGet:
				this.tmpMux.GET(v.UrlMapping(), WrapHttpApiHandler(v))
			case api.HttpMethodPost:
				this.tmpMux.POST(v.UrlMapping(), WrapHttpApiHandler(v))
			case api.HttpMethodDelete:
				this.tmpMux.DELETE(v.UrlMapping(), WrapHttpApiHandler(v))
			case api.HttpMethodPut:
				this.tmpMux.PUT(v.UrlMapping(), WrapHttpApiHandler(v))
			case api.HttpMethodPatch:
				this.tmpMux.PATCH(v.UrlMapping(), WrapHttpApiHandler(v))
			case api.HttpMethodHead:
				this.tmpMux.HEAD(v.UrlMapping(), WrapHttpApiHandler(v))
			case api.HttpMethodOptions:
				this.tmpMux.OPTIONS(v.UrlMapping(), WrapHttpApiHandler(v))
			default:
				panic(errors.New("unknown http method:" + fmt.Sprint(m)))
			}
		}
	}
}

func (this *HttpApi) Start() error {
	this.mux = this.tmpMux
	this.tmpMux = nil
	server := &http.Server{Addr: this.addr, Handler: this.mux}

	ln, err := net.Listen("tcp", this.addr)
	if err != nil {
		return err
	}

	this.server = server
	this.listener = ln

	go func() {
		err := server.Serve(ln)
		if err == http.ErrServerClosed {
			LOGGER.Info("http api is closed.")
		} else if err != nil {
			LOGGER.Error("serving from http api error.", err)
			panic(err)
		}
	}()

	LOGGER.Info("Http API endpoint starts at: <", this.addr, ">")

	return nil
}

func (this *HttpApi) ReloadHttpApi() error {
	this.mux = this.tmpMux
	this.tmpMux = nil
	this.server.Handler = this.mux
	return nil
}

func (this *HttpApi) Close() error {
	if err := this.listener.Close(); err != nil {
		return nil
	}
	if err := this.server.Shutdown(context.Background()); err != nil {
		return nil
	}
	return nil
}

func WrapHttpApiHandler(h api.HttpApiSimpleHandler) httprouter.Handle {
	reqType, resType := h.ContentTypes()
	t := h.RequestType()
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		switch request.Method {
		case http.MethodGet:
			fallthrough
		case http.MethodHead:
			fallthrough
		case http.MethodOptions:
			httpReq := api.HttpRequest{
				RequestURI: request.RequestURI,
				Params:     params,
			}
			//TODO MsbHandler required
			r := h.Handle(&httpReq, nil)
			writeResult(r, resType, writer)
			return
		}

		switch reqType {
		case api.RequestFormatJson:
			contentLength := request.ContentLength
			if contentLength == 0 {
				httpReq := api.HttpRequest{
					RequestURI: request.RequestURI,
					Params:     params,
				}
				//TODO MsbHandler required
				r := h.Handle(&httpReq, nil)
				writeResult(r, resType, writer)
				return
			} else {
				if !strings.HasPrefix(request.Header.Get("Content-Type"), _ContentTypeJson) {
					LOGGER.Error("content type is not acceptable:", request.Header.Get("Content-Type"))
					writeFailure(http.StatusInternalServerError, writer)
					return
				}
				data, err := io.ReadAll(request.Body)
				if err != nil {
					LOGGER.Error("read http body error:", err)
					writeFailure(http.StatusInternalServerError, writer)
					return
				}
				v := reflect.New(t)
				if err := utils.ObjFromJson(data, v.Interface()); err != nil {
					LOGGER.Error("unmarshal request from payload error:", err)
					writeFailure(http.StatusInternalServerError, writer)
					return
				}
				httpReq := api.HttpRequest{
					RequestURI:    request.RequestURI,
					Params:        params,
					RequestObject: v.Interface(),
				}
				//TODO MsbHandler required
				r := h.Handle(&httpReq, nil)
				writeResult(r, resType, writer)
				return
			}
		default:
			LOGGER.Error("unsupported request type:", reqType)
			writeFailure(http.StatusInternalServerError, writer)
			return
		}

	}
}

func writeResult(r *api.HttpResponse, resType api.ResponseFormat, writer http.ResponseWriter) {
	var data []byte

	switch resType {
	case api.ResponseFormatRawBinary:
		data = r.HandleResult.Result.([]byte)
		writer.Header().Set("Content-Type", r.HttpContentType)
	case api.ResponseFormatJson:
		data = utils.ObjToJson(r.HandleResult.Result)
		writer.Header().Set("Content-Type", _ContentTypeJson)
	}

	writer.WriteHeader(r.HttpStatus)
	_, err := writer.Write(data)
	if err != nil {
		LOGGER.Error("http api handler respond data error.", err)
	}
}

func writeFailure(status int, writer http.ResponseWriter) {
	writer.WriteHeader(status)
}
