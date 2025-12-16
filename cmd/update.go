/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	database "github/subham/CLI_TODO/Database"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "This id Update Functo",
	Long:  `We will pass an id the id of task and an string represents the Updated task name `,
	Args:  cobra.ExactArgs(1),
	Run:   update,
}

func init() {
	updateCmd.Flags().IntP("id", "i", 0, "This gives the id of the task to update")
	updateCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(updateCmd)
}

func update(cmd *cobra.Command, args []string) {
	value := args[0]
	_id, err := cmd.Flags().GetInt("id")
	if err != nil {
		fmt.Println("Error in getting flag")
	}
	fmt.Println("vlue ", value)
	fmt.Println("flag ", _id)
	database.Update(_id, value)
}
