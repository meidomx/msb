// +build !linux,!freebsd,!darwin !cgo

package pluginloader

import "github.com/meidomx/msb/config"

func InitPlugins(cfg config.MsbConfig) error {
	return nil
}
