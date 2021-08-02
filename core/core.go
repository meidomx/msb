package core

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/meidomx/msb/config"
	"github.com/meidomx/msb/kernel/httpapi"
	"github.com/meidomx/msb/module"
)

type MsbCore struct {
	httpApi *httpapi.HttpApi
	cfg     config.MsbConfig
}

func NewMsbCore(cfg config.MsbConfig) (*MsbCore, error) {
	msb := new(MsbCore)
	msb.httpApi = httpapi.NewHttpApi(cfg)

	return msb, nil
}

func (this *MsbCore) Init() error {
	//TODO load modules

	// init http api
	{
		hs := module.GetHttpApiHandlers()
		this.httpApi.BuildHttpApiHandlers(hs)
	}

	return nil
}

func (this *MsbCore) SyncStart() error {

	if err := this.httpApi.Start(); err != nil {
		return err
	}

	LOGGER.Info("MSB has started.")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	for s := range sigs {
		LOGGER.Debug("receive signal:", s)
		if err := this.Shutdown(); err != nil {
			LOGGER.Error("error occurs when shutting down MSB.", err)
		}
		break
	}

	return nil
}

func (this *MsbCore) Reload() error {
	//TODO
	return nil
}

func (this *MsbCore) Shutdown() error {
	//TODO
	return nil
}
