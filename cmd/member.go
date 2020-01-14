package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// memberCmd represents the member command
var memberCmd = &cobra.Command{
	Use:   "member",
	Short: "Add a member",
	Long:  `Will add a member to the family`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Provide only the name of the member")
		}
		name, err := validName(args[0])
		if err != nil {
			return err
		}

		return addMember(name)
	},
}

func init() {
	rootCmd.AddCommand(memberCmd)
}

func addMember(name string) error {
	if err := dbService.AddMember(name); err != nil {
		return err
	}

	fmt.Printf("Member %s successfully added\n", name)
	return nil
}
