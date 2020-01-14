package cmd

import (
	"fmt"
	"github.com/pkg/errors"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Counts the members with the given relation to given member",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("Provide only the following details in the given order - member, relation")
		}
		memberName, err := validName(args[0])
		if err != nil {
			return err
		}
		relation, err := validRelation(args[1])
		if err != nil {
			return err
		}

		relatives, err := getRelationList(memberName, relation)
		if err != nil {
			return errors.Errorf("Failed to get %s(S) of %s: %v", relation, memberName, err)
		}

		if len(relatives) == 0 {
			fmt.Printf("%s has no %sS\n", memberName, relation)
		} else {
			fmt.Printf("%s %s(S) are %v\n", memberName, relation, relatives)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func getRelationList(name string, relation string) ([]string, error) {
	list, err := dbService.ListRelations(name, relation)
	if err != nil {
		return nil, err
	}

	return list, nil
}
