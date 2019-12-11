package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/redemption/server/internal/config"
	"github.com/trustwallet/redemption/server/internal/storage"
	"github.com/trustwallet/redemption/server/platform"
	"os"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var (
	Storage *storage.Storage
	rootCmd = &cobra.Command{
		Use:   "Market Data",
		Short: "Sync all market data inside a database and provide an API to read",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			platform.Init()

			logger.Info("Connecting to the database...")
			var err error
			Storage, err = storage.New()
			if err != nil {
				logger.Fatal(err)
			}
			logger.Info("Database connected")
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("message", "m", "message.yml", "Message file")
	viper.BindPFlag("message.file", rootCmd.PersistentFlags().Lookup("message"))

	loadConf := func() { config.InitConfig() }
	loadLogger := func() { logger.InitLogger() }
	cobra.OnInitialize(loadConf)
	cobra.OnInitialize(loadLogger)
}
