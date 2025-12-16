/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	database "github/subham/CLI_TODO/Database"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "this is delete cmd ",
	Long:  `this is to delete by id give by user for flag -i`,
	Run:   Delete,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().IntP("id", "i", 0, "gives id as the flag")
	deleteCmd.MarkFlagRequired("id")
}
func Delete(cmd *cobra.Command, args []string) {
	id, err := cmd.Flags().GetInt("id")
	if err != nil {
		fmt.Println("Error getting Flag value")
		return
	}
	database.DeleteTask(id)
}
