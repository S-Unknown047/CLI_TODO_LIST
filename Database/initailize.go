package database

import (
	"fmt"
	"log"
	"os"
)

func Create() {
	fmt.Println("This is Inititazition")
	_, err := os.OpenFile("task.json", os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		if os.IsExist(err) {
			fmt.Println("File already exits")
			return
		}
		if os.IsPermission(err) {
			fmt.Println("Premission not found")
		}
		log.Fatal(err)
	}
	fmt.Println("File created")
}
