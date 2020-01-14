package cmd

import (
	"fmt"
	"github.com/pkg/errors"

	"github.com/spf13/cobra"
)

// countCmd represents the count command
var countCmd = &cobra.Command{
	Use:   "count",
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

		count, err := getRelationCount(memberName, relation)
		if err != nil {
			return errors.Errorf("Failed to get %s(S) of %s: %v", relation, memberName, err)
		}

		if count == 0 {
			fmt.Printf("%s has no %sS\n", memberName, relation)
		} else {
			fmt.Printf("%s has %d %s(S) \n", memberName, count, relation)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(countCmd)
}

func getRelationCount(name string, relation string) (int, error) {
	count, err := dbService.CountRelations(name, relation)
	if err != nil {
		return 0, err
	}

	return count, nil
}
