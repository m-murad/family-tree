/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/pkg/errors"

	"github.com/spf13/cobra"
)

// relationCmd represents the relation command
var relationCmd = &cobra.Command{
	Use:   "relation",
	Short: "Add a relation",
	Long:  `Will add a relation to the family`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Provide only the relation name")
		}

		relation, err := validRelation(args[0])
		if err != nil {
			return err
		}

		return addRelation(relation)
	},
}

func init() {
	rootCmd.AddCommand(relationCmd)
}

func addRelation(relation string) error {
	if err := dbService.AddRelation(relation); err != nil {
		return err
	}

	fmt.Printf("Relation %s successfully added\n", relation)
	return nil
}
