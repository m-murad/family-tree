package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all data",
	Long: `Will clear all the members and relationships in the database`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := clearData(); err != nil {
			return err
		}
		fmt.Println("All data has been cleared")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}

func clearData() error {
	if err := dbService.Clear(); err != nil {
		return err
	}

	return nil
}
