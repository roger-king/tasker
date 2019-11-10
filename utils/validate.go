package utils

import "encoding/json"

func Validate(args map[string]interface{}, p interface{}) error {
	jsonString, err := json.Marshal(args)

	if err != nil {
		return err
	}

	json.Unmarshal([]byte(jsonString), &p)

	return nil
}
