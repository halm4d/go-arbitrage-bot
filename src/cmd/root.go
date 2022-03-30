package cmd

import (
	"fmt"
	"github.com/halm4d/arbitragecli/app"
	"github.com/halm4d/arbitragecli/constants"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func readVersion() string {
	var version, err = os.ReadFile("src/.version")
	if err != nil {
		log.Fatalf("Cannot read version file. %s", err)
	}
	return string(version)
}

var rootCmd = &cobra.Command{
	Use:     "arbotgo",
	Version: readVersion(),
	Run: func(cmd *cobra.Command, args []string) {
		app.RunWebSocket()
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
	//rootCmd.PersistentFlags().BoolVarP(&constants.ShowOnlyProfitableArbs, "only-profitable", "p", false, "show only profitable arbitrages")
	rootCmd.PersistentFlags().Float64VarP(&constants.Fee, "fee", "f", .075, "fee")
	rootCmd.PersistentFlags().Float64VarP(&constants.BasePrice, "base-price", "b", 100, "base price")
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of arbotgo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Arbotgo version:  %s\n", rootCmd.Version)
	},
}
