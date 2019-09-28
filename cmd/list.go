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
	"os"

	"github.com/eloi/notary/pkg/notary"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

		var rpcServer, address, owner string
		var err error

		if rpcServer, err = cmd.Flags().GetString("rpcserver"); err != nil || len(rpcServer) == 0 {
			fmt.Println("--rpcserver not defined")
			os.Exit(1)
		}

		if address, err = cmd.Flags().GetString("address"); err != nil || len(address) == 0 {
			fmt.Println("--address not defined")
			os.Exit(1)
		}

		if owner, err = cmd.Flags().GetString("owner"); err != nil {
			fmt.Printf("%+v\n", err)
		}

		notary.InitClient(rpcServer, address)

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

	//listCmd.Flags().StringP("account", "x", viper.GetString("ACCOUNT"), "Ethereum account to operate from. Example: 0xFe5a....413f (required)")
	listCmd.Flags().StringP("address", "a", viper.GetString("ADDRESS"), "Smart contract address. Example: 0xFe5a....413f (required)")
	listCmd.Flags().StringP("rpcserver", "s", viper.GetString("RPCSERVER"), "Rpcserver where smart contract is deployed")
}
