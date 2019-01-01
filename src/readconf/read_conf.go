package readconf

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// ReadConfig : Read toml config file using viper
// author : Huripto Sugandi
// created date : 3 Dec 2018
func ReadConfig() *viper.Viper {
	v := viper.New()
	// set config file name (default : /<go root dir>/config/config.toml)
	v.SetConfigName("config")
	// set config file path (default : /<go root dir>/config)
	v.AddConfigPath("../../../config/")
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("couldn't load config: %s", err)
		os.Exit(1)
	}

	return v
}
