package config

import "github.com/spf13/viper"

func setServerDefault() {
	viper.SetDefault("server.port", 8000)
}

// SetDefault is a func will be called when init, injecting the defeat values of config
func SetDefault() {
	setServerDefault()
}
