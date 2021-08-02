package main

import (
	"os"

	"github.com/meidomx/msb/config"
	"github.com/meidomx/msb/core"
	"github.com/sirupsen/logrus"
)

import (
	_ "github.com/meidomx/msb/builtin"
)

var LOGGER *logrus.Logger

func init() {
	LOGGER = logrus.New()
	LOGGER.SetFormatter(&logrus.TextFormatter{})
}

func main() {
	LOGGER.Info("=======================================================================")
	LOGGER.Info("MMMMMMMM               MMMMMMMM   SSSSSSSSSSSSSSS BBBBBBBBBBBBBBBBB   ")
	LOGGER.Info("M:::::::M             M:::::::M SS:::::::::::::::SB::::::::::::::::B  ")
	LOGGER.Info("M::::::::M           M::::::::MS:::::SSSSSS::::::SB::::::BBBBBB:::::B ")
	LOGGER.Info("M:::::::::M         M:::::::::MS:::::S     SSSSSSSBB:::::B     B:::::B")
	LOGGER.Info("M::::::::::M       M::::::::::MS:::::S              B::::B     B:::::B")
	LOGGER.Info("M:::::::::::M     M:::::::::::MS:::::S              B::::B     B:::::B")
	LOGGER.Info("M:::::::M::::M   M::::M:::::::M S::::SSSS           B::::BBBBBB:::::B ")
	LOGGER.Info("M::::::M M::::M M::::M M::::::M  SS::::::SSSSS      B:::::::::::::BB  ")
	LOGGER.Info("M::::::M  M::::M::::M  M::::::M    SSS::::::::SS    B::::BBBBBB:::::B ")
	LOGGER.Info("M::::::M   M:::::::M   M::::::M       SSSSSS::::S   B::::B     B:::::B")
	LOGGER.Info("M::::::M    M:::::M    M::::::M            S:::::S  B::::B     B:::::B")
	LOGGER.Info("M::::::M     MMMMM     M::::::M            S:::::S  B::::B     B:::::B")
	LOGGER.Info("M::::::M               M::::::MSSSSSSS     S:::::SBB:::::BBBBBB::::::B")
	LOGGER.Info("M::::::M               M::::::MS::::::SSSSSS:::::SB:::::::::::::::::B ")
	LOGGER.Info("M::::::M               M::::::MS:::::::::::::::SS B::::::::::::::::B  ")
	LOGGER.Info("MMMMMMMM               MMMMMMMM SSSSSSSSSSSSSSS   BBBBBBBBBBBBBBBBB   ")
	LOGGER.Info("=======================================================================")

	cfg, err := config.LoadConfig("msb.toml")
	if err != nil {
		LOGGER.Error("load config file error.", err)
		os.Exit(1)
	}
	msb, err := core.NewMsbCore(*cfg)
	if err != nil {
		LOGGER.Error("new msb core failed.", err)
		os.Exit(1)
	}
	if err := msb.Init(); err != nil {
		LOGGER.Error("init msb core failed.", err)
		os.Exit(1)
	}
	if err := msb.SyncStart(); err != nil {
		LOGGER.Error("finish msb core failed.", err)
		os.Exit(1)
	}
}
