package pkg

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
