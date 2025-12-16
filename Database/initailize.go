package database

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func Create() {
	fmt.Println("This is Inititazition")
	os.OpenFile("task.json", os.O_CREATE|os.O_EXCL, 0644)
	f, err := os.ReadFile("task.json")

	if len(f) == 0 {

		js := []byte(`[
     	{
		   "count":1
		}
	]`)

		if !json.Valid(js) {
			fmt.Println("Error in jason ")
			return
		}
		fmt.Println("js :", string(js))
		errors := os.WriteFile("task.json", js, 0644)
		if errors != nil {
			log.Fatal(errors)
		}
	}

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
