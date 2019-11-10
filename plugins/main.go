package main

import (
	"fmt"

	"github.com/roger-king/tasker/utils"
)

type payload struct {
	Name string `json:"name" validate:"required"`
}

func Run(args map[string]interface{}) error {
	var p payload

	err := utils.Validate(args, p)

	if err != nil {
		return err
	}

	fmt.Println("Hello Tasker: ", p.Name)
	return nil
}
