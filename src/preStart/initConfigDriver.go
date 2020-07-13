package preStart

import (
	"github.com/a632079/ncm-helper/src/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

// The Config Parse Driver is served by viper
func initConfigDriver () {
	// Set default viper information
	viper.SetConfigName("config")
	viper.SetConfigType("toml") // Toml is the best!
	// Parse path etc > home > localPath
	viper.AddConfigPath("/etc/.ncm-helper")
	viper.AddConfigPath("$HOME/.ncm-helper")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../conf")
	if config.File != "" {
		viper.AddConfigPath(config.File)
	}
	// Parse env config
	viper.SetEnvPrefix("ncm") // like: NCM_PORT=8000
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("[preStart] Fatal error while reading config file: %s \n", err)
	}
}
