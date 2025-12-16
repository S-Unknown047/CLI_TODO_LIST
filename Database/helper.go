package database

import (
	"encoding/json"
	"fmt"
	"github/subham/CLI_TODO/model"
	"os"
	"time"
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

func Update(id int, description string) {
	data, err := os.ReadFile("task.json")
	if err != nil {
		fmt.Println("Error in Reading File")
		return
	}
	// Read file

	var arr []map[string]interface{}
	if err := json.Unmarshal(data, &arr); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	found := false
	for _, obj := range arr {
		if v, ok := obj["id"]; ok {
			fmt.Println("ID for traveral ", v, id)
			switch val := v.(type) {
			case float64:
				if int(val) == id {
					obj["description"] = description
					obj["updatedAt"] = time.Now().Format(time.RFC3339)
					found = true
				}
			case int:
				if val == id {
					obj["description"] = description
					obj["updatedAt"] = time.Now().Format(time.RFC3339)
					found = true
				}
			default:
			}
			if found {
				break
			}
		}
	}

	if !found {
		fmt.Println("Task id  not Found")
		return
	}
	filejson, err := json.MarshalIndent(arr, "", "\t")
	if err != nil {
		fmt.Println("Error in converting To json ", err)
		return
	}

	if json.Valid(filejson) {
		fmt.Println("Valid Json : \n", string(filejson))
		// persist changes
		if err := os.WriteFile("task.json", filejson, 0644); err != nil {
			fmt.Println("Error writing file:", err)
			return
		}
	} else {
		fmt.Println("Not valid json")
	}
}

func MarkedFunction(id int, status string) {
	data, err := os.ReadFile("task.json")
	if err != nil {
		fmt.Println("Error in reading file", err)
		return
	}

	var arr []map[string]interface{}
	err = json.Unmarshal(data, &arr)
	if err != nil {
		fmt.Println("Error in parsing json", err)
		return
	}
	found := false
	for _, obj := range arr {
		if v, ok := obj["id"]; ok {
			switch val := v.(type) {
			case float64:
				if int(val) == id {
					obj["Status"] = status
					obj["updateAt"] = time.Now()
					found = true
				}
			case int:
				if val == id {
					obj["Status"] = status
					obj["updateAt"] = time.Now()
					found = true
				}
			default:
			}

			if found {
				break
			}
		}
	}
	if !found {
		fmt.Println("Not found")
		return
	}

	jsonFile, err := json.MarshalIndent(arr, "", "\t")

	if err != nil {
		fmt.Println("Error in Unmarshall")
		return
	}

	if json.Valid(jsonFile) {
		os.WriteFile("task.json", jsonFile, 0644)
	}

	fmt.Println("Invalid json")
}

func DeleteTask(id int) {
	data, err := os.ReadFile("task.json")
	if err != nil {
		fmt.Println("Error while Reading the file")
	}

	var arr []map[string]interface{}
	json.Unmarshal(data, &arr)
	var ind int = -1
	for index, obj := range arr {
		if v, ok := obj["id"]; ok {
			switch val := v.(type) {
			case float64:
				if int(val) == id {
					ind = index
				}
			case int:
				if int(val) == id {
					ind = index
				}
			}
			if ind != -1 {
				break
			}
		}
	}
	if ind == -1 {
		fmt.Println("Not found")
		return
	}
	deleteTask := arr[ind : ind+1]
	arr = append(arr[:ind], arr[ind+1:]...)

	file, err := json.MarshalIndent(arr, "", "\t")
	if err != nil {
		fmt.Println("Error in converting to json")
	}
	if !json.Valid(file) {
		fmt.Println("Json file is not valid")
		return
	}
	err = os.WriteFile("task.json", file, 0644)
	if err != nil {
		fmt.Println("Unsuccesful file writing")
		return
	}
	fmt.Println("Delete Sucessful", deleteTask)
}
