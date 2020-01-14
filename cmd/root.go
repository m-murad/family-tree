package cmd

import (
	"fmt"
	"github.com/m-murad/family-tree/storage"
	"github.com/spf13/cobra"
	"os"
)

var dbService storage.Service

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "family-tree",
	Short: "Define and query a family",
	Long: `This program lets you to create a family. 
You can add members to the family and define relationships between members.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	var err error
	dbService, err = storage.Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to database: %v\n", err)
		os.Exit(2)
	}
}
