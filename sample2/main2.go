package sample2

import (
	"encoding/json"
	"fmt"
	"time"
)

func Main2() {
	data := map[string]interface{}{
		"intValue":    1234,
		"boolValue":   true,
		"stringValue": "hello!",
		"dateValue":   time.Date(2023, 3, 3, 8, 18, 0, 0, time.UTC),
		"objectValue": map[string]interface{}{
			"arrayValue": []int{1, 2, 3, 4},
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	fmt.Printf("json data: %s\n", jsonData)
}

// https://www.digitalocean.com/community/tutorials/how-to-use-json-in-go
