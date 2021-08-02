package config

import (
	"github.com/meidomx/msb/payload/toml"
)

type MsbConfig struct {
	Global struct {
		HttpAddr     string `toml:"http_addr"`
		HttpsAddr    string `toml:"https_addr"`
		HttpApiAddr  string `toml:"http_api_addr"`
		HttpsApiAddr string `toml:"https_api_addr"`
	} `toml:"global"`
}

func LoadConfig(path string) (*MsbConfig, error) {
	config := new(MsbConfig)
	err := toml.LoadObjectFromFilePath(path, config)
	if err != nil {
		return nil, err
	} else {
		return config, nil
	}
}
