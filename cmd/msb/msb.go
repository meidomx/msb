package msb

import (
	"os"

	"github.com/meidomx/msb/config"
	"github.com/meidomx/msb/core"
	"github.com/sirupsen/logrus"
)

var LOGGER_MAIN *logrus.Logger

func init() {
	LOGGER_MAIN = logrus.New()
	LOGGER_MAIN.SetFormatter(&logrus.TextFormatter{})
}

func Run() {
	LOGGER_MAIN.Info("=======================================================================")
	LOGGER_MAIN.Info("MMMMMMMM               MMMMMMMM   SSSSSSSSSSSSSSS BBBBBBBBBBBBBBBBB   ")
	LOGGER_MAIN.Info("M:::::::M             M:::::::M SS:::::::::::::::SB::::::::::::::::B  ")
	LOGGER_MAIN.Info("M::::::::M           M::::::::MS:::::SSSSSS::::::SB::::::BBBBBB:::::B ")
	LOGGER_MAIN.Info("M:::::::::M         M:::::::::MS:::::S     SSSSSSSBB:::::B     B:::::B")
	LOGGER_MAIN.Info("M::::::::::M       M::::::::::MS:::::S              B::::B     B:::::B")
	LOGGER_MAIN.Info("M:::::::::::M     M:::::::::::MS:::::S              B::::B     B:::::B")
	LOGGER_MAIN.Info("M:::::::M::::M   M::::M:::::::M S::::SSSS           B::::BBBBBB:::::B ")
	LOGGER_MAIN.Info("M::::::M M::::M M::::M M::::::M  SS::::::SSSSS      B:::::::::::::BB  ")
	LOGGER_MAIN.Info("M::::::M  M::::M::::M  M::::::M    SSS::::::::SS    B::::BBBBBB:::::B ")
	LOGGER_MAIN.Info("M::::::M   M:::::::M   M::::::M       SSSSSS::::S   B::::B     B:::::B")
	LOGGER_MAIN.Info("M::::::M    M:::::M    M::::::M            S:::::S  B::::B     B:::::B")
	LOGGER_MAIN.Info("M::::::M     MMMMM     M::::::M            S:::::S  B::::B     B:::::B")
	LOGGER_MAIN.Info("M::::::M               M::::::MSSSSSSS     S:::::SBB:::::BBBBBB::::::B")
	LOGGER_MAIN.Info("M::::::M               M::::::MS::::::SSSSSS:::::SB:::::::::::::::::B ")
	LOGGER_MAIN.Info("M::::::M               M::::::MS:::::::::::::::SS B::::::::::::::::B  ")
	LOGGER_MAIN.Info("MMMMMMMM               MMMMMMMM SSSSSSSSSSSSSSS   BBBBBBBBBBBBBBBBB   ")
	LOGGER_MAIN.Info("=======================================================================")

	cfg, err := config.LoadConfig("msb.toml")
	if err != nil {
		LOGGER_MAIN.Error("load config file error.", err)
		os.Exit(1)
	}
	msb, err := core.NewMsbCore(*cfg)
	if err != nil {
		LOGGER_MAIN.Error("new msb core failed.", err)
		os.Exit(1)
	}
	if err := msb.Init(); err != nil {
		LOGGER_MAIN.Error("init msb core failed.", err)
		os.Exit(1)
	}
	if err := msb.SyncStart(); err != nil {
		LOGGER_MAIN.Error("finish msb core failed.", err)
		os.Exit(1)
	}
}
