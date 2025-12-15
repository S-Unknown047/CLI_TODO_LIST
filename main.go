/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	database "github/subham/CLI_TODO/Database"
	"github/subham/CLI_TODO/cmd"
)

func main() {
	fmt.Println("This is CLI app")
	fmt.Println("This is File Creation")
	database.Create()
	cmd.Execute()
}
