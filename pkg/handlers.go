package pkg

import (
	"encoding/json"
	"net/http"

	db "upper.io/db.v3"
)

// ListTasks -
func ListTasks(session db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer session.Close()
		taskService := NewTaskService(session)
		tasks, err := taskService.List()

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "", err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, tasks)
		return
	}
}

// Response Helpers:
type errorHelper struct {
	Error       string `json:"error"`
	Description string `json:"description"`
}

type dataHelper struct {
	Data interface{} `json:"data"`
}

func respondWithError(w http.ResponseWriter, code int, errorType string, message string) {
	response, _ := json.Marshal(&errorHelper{Error: errorType, Description: message})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(&dataHelper{Data: payload})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
