package cmd

import (
	"github.com/spf13/cobra"
	"github.com/trustwallet/redemption/server/internal/api"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "provide a redemption API",
	Run:   provideApi,
}

func provideApi(_ *cobra.Command, args []string) {
	api.Provide(Storage)
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
