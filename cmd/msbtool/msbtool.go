package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	cmd := flag.Arg(0)
	cmd = strings.ToLower(cmd)

	if cmd == "" {
		cmd = "run"
	}

	switch cmd {
	case "run":
		runCmd()
	default:
		fmt.Println("unknown command:" + cmd)
		os.Exit(1)
	}
}

func runCmd() {
	f, err := os.Create("msb.temp.go")
	if err != nil {
		panic(err)
	}
	defer func() {
		f.Close()
		os.Remove("msb.temp.go")
	}()
	cf, err := os.Create("msb.toml")
	if err != nil {
		panic(err)
	}
	defer func() {
		cf.Close()
		os.Remove("msb.toml")
	}()
	_, err = f.WriteString(mainFile)
	if err != nil {
		panic(err)
	}
	if err = f.Sync(); err != nil {
		panic(err)
	}
	_, err = cf.WriteString(configFile)
	if err != nil {
		panic(err)
	}
	if err = cf.Sync(); err != nil {
		panic(err)
	}

	if err = runBinary("go", []string{"get", "-u"}); err != nil {
		panic(err)
	}
	if err = runBinary("go", []string{"build", "-o", "a.out"}); err != nil {
		panic(err)
	}
	defer func() {
		os.Remove("a.out")
	}()
	if err = runBinary("a.out", nil); err != nil {
		panic(err)
	}
}

func runBinary(binaryName string, args []string) error {

	cmd := exec.Command(binaryName, args...)

	out, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	errOut, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for s := range signalChan {
			switch s {
			case os.Interrupt:
				cmd.Process.Signal(os.Interrupt)
			}
		}
	}()

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := out.Read(buf)
			if n > 0 {
				fmt.Fprint(os.Stdout, string(buf[:n]))
			}
			if err != nil {
				if err == io.EOF {
					return
				} else {
					panic(err)
				}
			}
		}
	}()
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := errOut.Read(buf)
			if n > 0 {
				fmt.Fprint(os.Stderr, string(buf[:n]))
			}
			if err != nil {
				if err == io.EOF {
					return
				} else {
					panic(err)
				}
			}
		}
	}()

	err = cmd.Start()
	if err != nil {
		return err
	}
	fmt.Println("starting service... pid:" + strconv.Itoa(cmd.Process.Pid))
	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}

const configFile = `
[global]
http_addr = ":8080"
https_addr = ":8443"
http_api_addr = ":8088"
https_api_addr = ":8488"

[shared.jobscheduling]
timezone = "Asia/Shanghai"

[plugins]
plugin_libs = []
plugin_folders = []
`

const mainFile = `
package main

import (
	"os"

	"github.com/meidomx/msb/config"
	"github.com/meidomx/msb/core"
	"github.com/sirupsen/logrus"
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
`
