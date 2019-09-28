/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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

	"github.com/eloi/notary/pkg/notary"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List registered assets",
	Long: `
	Use list to show all or part of the assets registered
	in the notary dapp. For example:

	notary list              list all registered assets
	notary list -o owner     list assets registered by "owner"
	`,
	Run: func(cmd *cobra.Command, args []string) {

		owner, err := cmd.Flags().GetString("owner")
		if err != nil {
			fmt.Printf("%+v\n", err)
		}

		if len(owner) > 0 {
			notary.ListByOwner(owner)
		} else {
			notary.ListAll()
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	listCmd.Flags().StringP("owner", "o", "", "Show only assets of this owner")
}
