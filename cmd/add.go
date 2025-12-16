/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	database "github/subham/CLI_TODO/Database"
	"github/subham/CLI_TODO/model"
	"log"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "ADD A TASK TO TODO LIST",
	Long:  `ADD A TASK TO YOUR TO-DO list`,
	Args:  cobra.ExactArgs(1),
	Run:   AddCmd,
}

func init() {

	addCmd.Flags().StringP(
		"status",
		"f",
		"todo",
		"This is to define whats the status of the task")
	rootCmd.AddCommand(addCmd)

}
func AddCmd(cmd *cobra.Command, args []string) {
	value := args[0]
	status, err := cmd.Flags().GetString("status")
	if err != nil {
		fmt.Println("Error in getting flag")
		return
	}
	fmt.Println("I am in AddCMD")
	fmt.Println("Value", value)
	fmt.Println("status", status)

	newStatus := strings.TrimSpace(strings.ToLower(status))

	if newStatus != "todo" && newStatus != "done" && newStatus != "in-progress" {
		fmt.Printf("Entered Wrong Flag Value Flag Values can be only\n 1) done  if task is complete \n2) todo if task is not yet started\n3) in-progress if task is in progress")
	}

	structValue := model.Task{Description: value, Status: newStatus, CreatedAt: time.Now(), UpdatedAt: time.Now()}

	fmt.Println("Calling Add function")
	err = database.Add(structValue)
	if err != nil {
		fmt.Println("Error while adding Task")
		log.Fatal(err)
	}
}
