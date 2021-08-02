package module

import "github.com/meidomx/msb/api"

var hs []api.HttpApiHandler

func RegisterHttpApiHandler(h api.HttpApiHandler) {
	hs = append(hs, h)
}

func GetHttpApiHandlers() []api.HttpApiHandler {
	return hs
}
