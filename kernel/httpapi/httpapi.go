package httpapi

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/meidomx/msb/api"
	"github.com/meidomx/msb/config"

	"github.com/julienschmidt/httprouter"
)

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

func (this *HttpApi) BuildHttpApiHandlers(hs []api.HttpApiHandler) {
	this.tmpMux = httprouter.New()
	for _, v := range hs {
		for _, m := range v.HttpMethods() {
			switch m {
			case api.HttpMethodGet:
				this.tmpMux.GET(v.UrlMapping(), v.Handler())
			case api.HttpMethodPost:
				this.tmpMux.POST(v.UrlMapping(), v.Handler())
			case api.HttpMethodDelete:
				this.tmpMux.DELETE(v.UrlMapping(), v.Handler())
			case api.HttpMethodPut:
				this.tmpMux.PUT(v.UrlMapping(), v.Handler())
			case api.HttpMethodPatch:
				this.tmpMux.PATCH(v.UrlMapping(), v.Handler())
			case api.HttpMethodHead:
				this.tmpMux.HEAD(v.UrlMapping(), v.Handler())
			case api.HttpMethodOptions:
				this.tmpMux.OPTIONS(v.UrlMapping(), v.Handler())
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
