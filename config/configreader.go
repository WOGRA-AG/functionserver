package config

import (
	"log"

	"github.com/spf13/viper"
)

type RestConfig struct {
	Host           string
	Port           string
	TrustedProxies []string
}

type DbConfig struct {
	DatabaseName string
	DatabaseUrl  string
	DatabasePort string
	Login        string
	Password     string
}

func ReadRestConfiguration() RestConfig {

	// Set the file name of the configurations file
	viper.SetConfigName("rest")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Unable to decode into struct, %v", err)
	}

	var restConf RestConfig

	err := viper.Unmarshal(&restConf)
	if err != nil {
		log.Printf("Unable to decode into struct, %v", err)
	}

	return restConf
}

func ReadDatabaseConfiguration() DbConfig {

	// Set the file name of the configurations file
	viper.SetConfigName("db")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Unable to decode into struct, %v", err)
	}

	var dbConf DbConfig

	err := viper.Unmarshal(&dbConf)
	if err != nil {
		log.Printf("Unable to decode into struct, %v", err)
	}

	return dbConf
}
