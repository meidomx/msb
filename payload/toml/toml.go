package toml

import (
	"github.com/BurntSushi/toml"
)

func LoadObjectFromFilePath(path string, obj interface{}) error {
	_, err := toml.DecodeFile(path, obj)
	return err
}
