package config

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type configuration struct {
	Port string
	Code struct {
		Prefix  string
		Pattern string
	}
	Link struct {
		Url string
	}
	Wallet struct {
		Mnemonic string
	}
	Mongo struct {
		Uri string
	}
	Api struct {
		Mode       string
		Auth_Token string
	}
	Message struct {
		File string
	}
}

var Configuration configuration

// set dummy values to force viper to search for these keys in environment variables
// the AutomaticEnv() only searches for already defined keys in a config file, default values or kvstore struct.
func setDefaults() {
	viper.SetDefault("Port", "8399")
	viper.SetDefault("Code.Prefix", "")
	viper.SetDefault("Code.Pattern", "####-####-####")
	viper.SetDefault("Link.Url", "https://links.trustwallet.com/redeem?code={{.Code}}&provider={{.Provider}}")
	viper.SetDefault("Wallet.Mnemonic", "")
	viper.SetDefault("Mongo.Uri", "")
	viper.SetDefault("Api.Mode", "release")
	viper.SetDefault("Api.Auth_Token", "")
	viper.SetDefault("Message.File", "")
}

// initConfig reads in config file and ENV variables if set.
func InitConfig() {
	setDefaults()
	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.Unmarshal(&Configuration); err != nil {
		fmt.Printf("Error Unmarshal: %s \n", err)
	}
	log.Printf("API_PORT: %s", Configuration.Port)
	log.Printf("CODE_PREFIX: %s", Configuration.Code.Prefix)
	log.Printf("CODE_PATTERN: %s", Configuration.Code.Pattern)
	log.Printf("LINK_URL: %s", Configuration.Link.Url)
	log.Printf("MONGO_URI: %s", Configuration.Mongo.Uri)
	log.Printf("API_MODE: %s", Configuration.Api.Mode)
	log.Printf("API_AUTH_TOKEN: %s", Configuration.Api.Auth_Token)
	log.Printf("MESSAGE_FILE: %s", Configuration.Message.File)

	mnemonic := ""
	if len(Configuration.Wallet.Mnemonic) > 0 {
		mnemonic = "*********************"
	}
	log.Printf("WALLET_MNEMONIC: %s", mnemonic)
}
