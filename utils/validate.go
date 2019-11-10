package utils

import (
	"encoding/json"

	"gopkg.in/go-playground/validator.v9"
)

func Validate(args map[string]interface{}, p interface{}) error {
	var err error

	v := validator.New()
	jsonString, err := json.Marshal(args)

	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(jsonString), &p)

	if err != nil {
		return err
	}

	err = v.Struct(p)

	if err != nil {
		return err
	}

	return nil
}
