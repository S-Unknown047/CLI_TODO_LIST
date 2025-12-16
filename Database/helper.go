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
	count := 0

	// find "count":<number> in ReadFile, parse it, increment and replace in ReadFile
	count = 0
	for i := 0; i+7 < len(ReadFile); i++ {
		if string(ReadFile[i:i+7]) == "\"count\"" {
			// find the colon after "count"
			j := i + 7
			for j < len(ReadFile) && ReadFile[j] != ':' {
				j++
			}
			if j >= len(ReadFile) {
				continue
			}
			j++ // move past ':'
			// skip whitespace
			for j < len(ReadFile) && (ReadFile[j] == ' ' || ReadFile[j] == '\t' || ReadFile[j] == '\n' || ReadFile[j] == '\r') {
				j++
			}
			// parse consecutive digits
			start := j
			for j < len(ReadFile) && ReadFile[j] >= '0' && ReadFile[j] <= '9' {
				count = count*10 + int(ReadFile[j]-'0')
				j++
			}
			end := j
			if start == end {
				continue
			}
			// increment count and replace the number bytes in ReadFile
			newCount := count + 1
			newNum := []byte(fmt.Sprintf("%d", newCount))
			ReadFile = append(ReadFile[:start], append(newNum, ReadFile[end:]...)...)
			break
		}
	}

	os.WriteFile("task.json", ReadFile, 0644)
	file, err := os.OpenFile("task.json",
		os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	task.Id = count
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
