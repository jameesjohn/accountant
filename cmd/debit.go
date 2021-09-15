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

// debitCmd represents the debit command
var debitCmd = &cobra.Command{
	Use:   "debit",
	Short: "Create a debit transaction",
	Long: `
This command creates a debit transaction for a particular user.
Usage: accountant debit <username> --amount=<amount> --narration=<narration>.`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Username not specified")
		}
		username := args[0]

		user, err := database.FindOrCreateUser(username)
		if err != nil {
			log.Fatal(err)
		}
		if user.Balance > debitAmount {
			user.Balance = user.Balance - debitAmount
			debitTransaction := database.Transaction{Amount: debitAmount, Type: "debit", Narration: debitNarration}
			user.Transactions = append(user.Transactions, debitTransaction)
			database.UpdateUser(user)

			fmt.Println("Transaction created successfully")
		} else {
			fmt.Println("Insufficient funds!")
		}

	},
}

var debitNarration string
var debitAmount int64

func init() {
	rootCmd.AddCommand(debitCmd)

	debitCmd.Flags().StringVarP(&debitNarration, "narration", "n", "", "Narration for this debit transaction")
	debitCmd.Flags().Int64VarP(&debitAmount, "amount", "a", 0, "Amount to be debited")

	debitCmd.MarkFlagRequired("narration")
	debitCmd.MarkFlagRequired("amount")
}
