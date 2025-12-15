package database

import (
	"encoding/json"
	"fmt"
	"github/subham/CLI_TODO/model"
	"os"
)

func Add(task model.Task) error {
	fmt.Println("in Add Function")
	ReadFile, _ := os.ReadFile("task.json")
	ReadFile[len(ReadFile)-2] = byte(',')
	ReadFile[len(ReadFile)-1] = byte(' ')
	os.WriteFile("task.json", ReadFile, 0644)
	file, err := os.OpenFile("task.json",
		os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	jsonFile, err := json.MarshalIndent(task, "", "\t")
	if err != nil {
		return err
	}

	fmt.Println("Json", string(jsonFile))
	_, err =
		file.Write(jsonFile)
	file.Write([]byte("\n]"))

	if err != nil {
		return err
	}
	return nil
}

// func Update(id int) error{
// 	data, err := os.ReadFile("../task.json")
// 	if err != nil {
// 		fmt.Println("Error in Reading File")
// 		return err
// 	}
// 	var val map[string]interface{}
// 	valid := json.Valid(data)
//     if valid{
// 		err:=json.Unmarshal(data,&val)
// 		if err!=nil{
// 			fmt.Println("Error in converting File ",err)
// 			return err
// 		}
// 	}
// }

// func Delete() {

// }
