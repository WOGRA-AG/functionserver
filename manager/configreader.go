package manager

import (
	"log"

	"github.com/spf13/viper"
)

type AccessTokenList struct {
	Accesstokens []AccessToken `yaml:"Accesstokens"`
}

type AccessToken struct {
	Token    string `json:"accesstoken" yaml:"token"`
	User     string `json:"user"  yaml:"user"`
	Password string `json:"password"  yaml:"password"`
}

type RestConfig struct {
	Host string
	Port string
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

func ReadAccessTokensConfiguration() []AccessToken {

	// Set the file name of the configurations file
	viper.SetConfigName("accesstoken")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Unable to decode into struct, %v", err)
	}

	viper.Get("name")

	var access AccessTokenList

	err := viper.Unmarshal(&access)
	if err != nil {
		log.Printf("Unable to decode into struct, %v", err)
	}

	log.Printf("Tokens read: %v", access)
	return access.Accesstokens
}

func CheckAccessToken(token *AccessToken) bool {

	if token == nil {
		return false
	}
	// at the moment only the token id will be checked
	// user and password will be ignored currently
	return HasAccessToken(token.Token)
}

func HasAccessToken(token string) bool {

	if token == "" {
		log.Println("Token is empty.")
		return false
	}

	tokens := ReadAccessTokensConfiguration()
	return contains(tokens, token)
}

func contains(s []AccessToken, e string) bool {
	for _, a := range s {
		if a.Token == e {
			log.Println("Given AccesToken found.")
			return true
		}
	}

	log.Println("Given AccesToken not found.")
	return false
}
