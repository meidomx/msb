package module

import "github.com/meidomx/msb/api"

var hs []api.HttpApiSimpleHandler

func RegisterHttpApiHandler(h api.HttpApiSimpleHandler) {
	hs = append(hs, h)
}

func GetHttpApiHandlers() []api.HttpApiSimpleHandler {
	return hs
}
