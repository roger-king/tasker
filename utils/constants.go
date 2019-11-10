package utils

import (
	"os"
)

var (
	// DBType -
	DBType = os.Getenv("TASKER_DB_TYPE")

	// DBHost -
	DBHost = os.Getenv("TASKER_DB_HOST")

	// DBUser -
	DBUser = os.Getenv("TASKER_DB_USER")

	// DBPassword -
	DBPassword = os.Getenv("TASKER_DB_PASSWORD")

	// DBPort -
	DBPort = os.Getenv("TASKER_DB_PORT")

	// ScriptRoot -
	ScriptRoot = os.Getenv("TASKER_SCRIPT_ROOT")

	// TASKER_ENV -
	TaskerEnv = os.Getenv("TASKER_ENV")
)

// HTTPError -
type HTTPError string

func (h HTTPError) String() string {
	return string(h)
}

const (
	// ProcessingError -
	ProcessingError HTTPError = "processing_error"
	// RequestError -
	RequestError HTTPError = "request_error"
)