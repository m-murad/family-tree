package cmd

import (
	"fmt"
	"github.com/pkg/errors"

	"github.com/spf13/cobra"
)

// relateCmd represents the relate command
var relateCmd = &cobra.Command{
	Use:   "relate",
	Short: "Relate to members",
	Long:  `Will create a relationship between members of a family`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 3 {
			return errors.New("Provide only the following details in the given order - firstMember, relation, secondNumber")
		}

		firstMember, err := validName(args[0])
		if err != nil {
			return err
		}

		relation, err := validRelation(args[1])
		if err != nil {
			return err
		}

		secondMember, err := validName(args[2])
		if err != nil {
			return err
		}

		if firstMember == secondMember {
			return errors.New("Please specify 2 different members")
		}

		return createRelationship(firstMember, relation, secondMember)
	},
}

func init() {
	rootCmd.AddCommand(relateCmd)
}

func createRelationship(firstMember string, relation string, secondMember string) error {
	if err := dbService.DefineRelation(firstMember, relation, secondMember); err != nil {
		return err
	}

	fmt.Println("Relationship defined successfully")
	return nil
}
