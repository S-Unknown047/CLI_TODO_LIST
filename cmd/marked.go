/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	database "github/subham/CLI_TODO/Database"

	"github.com/spf13/cobra"
)

// markedCmd represents the marked command
var markedCmd = &cobra.Command{
	Use:   "marked",
	Short: "It mark the todo , done , in-progress",
	Long:  `it have a flag with -i for id and a argument in string which havve done, todo, in-progress.`,
	Args:  cobra.ExactArgs(1),
	Run:   marked,
}

func init() {
	rootCmd.AddCommand(markedCmd)
	markedCmd.Flags().IntP("id", "i", 0, "This is id of todo task")
	markedCmd.MarkFlagRequired("id")
}

func marked(cmd *cobra.Command, args []string) {
	val := args[0]
	id, _ := cmd.Flags().GetInt("id")

	fmt.Println("marked called")
	fmt.Println("Value ", val)
	fmt.Println("Flag \n\n", id)

	database.MarkedFunction(id, val)

}
