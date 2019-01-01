package read_conf

	import (
		"fmt"
		"github.com/spf13/viper"
		"os"
	)

	func ReadConfig()(*viper.Viper) {
		v := viper.New()
		v.SetConfigName("config")
		v.AddConfigPath("../config/")
		if err := v.ReadInConfig(); err != nil {
			fmt.Printf("couldn't load config: %s", err)
			os.Exit(1)
		}

		return v
	}

