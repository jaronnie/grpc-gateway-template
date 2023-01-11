package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"github.com/jaronnie/grpc-gateway-template/base/util"
)

const (
	defaultFilePath = "./conf/cfg.toml"
)

// NewLocalConfig return local configure instance
func NewLocalConfig() (err error) {
	// set local file path
	viper.SetConfigFile(defaultFilePath)
	ext, err := util.Ext(defaultFilePath, viper.SupportedExts...)
	if err != nil {
		panic(err)
	}
	viper.SetConfigType(ext)

	if err = viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// watch config and set onchange function
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// TODO: local file change
		fmt.Println(fmt.Sprintf("Config file changed: %s\n", e.Name))
	})
	return
}
