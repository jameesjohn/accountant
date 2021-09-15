/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"log"

	"github.com/jameesjohn/accountant/database"
	"github.com/spf13/cobra"
)

// creditCmd represents the credit command
var creditCmd = &cobra.Command{
	Use:   "credit",
	Short: "Create a credit transaction",
	Long: `
This command creates a credit transaction for a particular user.
Usage: accountant credit <username> --amount=<amount> --narration=<narration>.`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Username not specified")
		}
		username := args[0]

		user, err := database.FindOrCreateUser(username)
		if err != nil {
			log.Fatal(err)
		}

		user.Balance = user.Balance + creditAmount
		creditTransaction := database.Transaction{Amount: creditAmount, Type: "credit", Narration: creditNarration}
		user.Transactions = append(user.Transactions, creditTransaction)
		database.UpdateUser(user)

		fmt.Println("Transaction created successfully")

	},
}

var creditNarration string
var creditAmount int64

func init() {
	rootCmd.AddCommand(creditCmd)

	creditCmd.Flags().StringVarP(&creditNarration, "narration", "n", "", "Narration for this credit transaction")
	creditCmd.Flags().Int64VarP(&creditAmount, "amount", "a", 0, "Amount to be credited")

	creditCmd.MarkFlagRequired("narration")
	creditCmd.MarkFlagRequired("amount")
}
