package main

import (
	"encoding/json"
	"fmt"
)

type payload struct {
	Name string `json:"name"`
}

func Run(args map[string]interface{}) error {
	var p payload
	jsonString, err := json.Marshal(args)

	if err != nil {
		return err
	}

	json.Unmarshal([]byte(jsonString), &p)

	fmt.Println("Hello Tasker: ", p.Name)
	return nil
}
