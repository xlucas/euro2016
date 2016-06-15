package util

import "github.com/spf13/viper"

func LoadConfig() error {
	// Default
	viper.SetDefault("token", "")
	viper.SetDefault("emoji", false)

	// Config path
	viper.AddConfigPath("$HOME")
	viper.SetConfigName(".euro2016")
	viper.SetConfigType("json")

	// Read config
	return viper.ReadInConfig()
}
