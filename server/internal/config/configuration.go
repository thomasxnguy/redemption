package config

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"log"
	"strings"
)

type configuration struct {
	Port string
	Code struct {
		Prefix  string
		Pattern string
	}
	Client struct {
		Path string
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
	Transaction struct {
		Memo string
	}
}

var Configuration configuration

// set dummy values to force viper to search for these keys in environment variables
// the AutomaticEnv() only searches for already defined keys in a config file, default values or kvstore struct.
func setDefaults() {
	viper.SetDefault("Port", "8399")
	viper.SetDefault("Code.Prefix", "")
	viper.SetDefault("Code.Pattern", "####-####-####")
	viper.SetDefault("Client.Path", "./../client/build")
	viper.SetDefault("Link.Url", "https://link.trustwallet.com/redeem?code={{.Code}}&provider={{.Provider}}")
	viper.SetDefault("Wallet.Mnemonic", "")
	viper.SetDefault("Mongo.Uri", "")
	viper.SetDefault("Api.Mode", "release")
	viper.SetDefault("Api.Auth_Token", "")
	viper.SetDefault("Message.File", "")
	viper.SetDefault("Transaction.Memo", "Trust Wallet Redeem")
}

// initConfig reads in config file and ENV variables if set.
func InitConfig() {
	setDefaults()
	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.Unmarshal(&Configuration); err != nil {
		logger.Error(err, "Error Unmarshal Viper Config File")
	}
	log.Printf("API_PORT: %s", Configuration.Port)
	log.Printf("CODE_PREFIX: %s", Configuration.Code.Prefix)
	log.Printf("CODE_PATTERN: %s", Configuration.Code.Pattern)
	log.Printf("CLIENT_PATH: %s", Configuration.Client.Path)
	log.Printf("LINK_URL: %s", Configuration.Link.Url)
	log.Printf("WALLET_MNEMONIC: %s", Configuration.Wallet.Mnemonic)
	log.Printf("MONGO_URI: %s", Configuration.Mongo.Uri)
	log.Printf("API_MODE: %s", Configuration.Api.Mode)
	log.Printf("API_AUTH_TOKEN: %s", Configuration.Api.Auth_Token)
	log.Printf("MESSAGE_FILE: %s", Configuration.Message.File)
	log.Printf("TRANSACTION_MEMO: %s", Configuration.Transaction.Memo)
}
