/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	database "github/subham/CLI_TODO/Database"
	"strings"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the task based of value",
	Long: `four values : all ,
	todo, in-progress, done`,
	Args: cobra.ExactArgs(1),
	Run:  list,
}

func init() {
	rootCmd.AddCommand(listCmd)

}

func list(cmd *cobra.Command, args []string) {
	val := args[0]
	val = strings.ToLower(val)
	if val != "all" && val != "done" && val != "in-progress" && val != "todo" {
		fmt.Println("Val incorrect")
		return
	}
	database.Listing(val)
}
