package models

type Payload interface {
	Run(args map[string]interface{}) error
}