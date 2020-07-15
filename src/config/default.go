package config

import "github.com/spf13/viper"

func setServerDefault() {
	viper.SetDefault("server.port", 8000)
}

func SetDefault() {
	setServerDefault()
}
