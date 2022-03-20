package cmd

import (
	"fmt"
	"github.com/halm4d/arbitragecli/app"
	"github.com/halm4d/arbitragecli/constants"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:     "arbcli",
	Version: "v0.0.1-alpha",
	Run: func(cmd *cobra.Command, args []string) {
		app.Run()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().BoolVarP(&constants.Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().Float64VarP(&constants.Fee, "fee", "f", .075, "fee")
	rootCmd.PersistentFlags().Float64VarP(&constants.BasePrice, "base-price", "b", 100, "base price")
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of ArbCli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Arbitragecli  %s\n", rootCmd.Version)
	},
}
