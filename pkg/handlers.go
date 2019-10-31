package pkg

import (
	"encoding/json"
	"io"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
	"github.com/robfig/cron"
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

// CreateTask -
func CreateTask(session db.Database, scheduler *cron.Cron) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input NewInputTask

		defer session.Close()

		_, err := jsonDecode(input, r.Body)

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, RequestError, err.Error())
			return
		}

		scheduler.AddFunc(input.Schedule, func() {
			exec.Command(input.Executor)
		})

		taskService := NewTaskService(session)
		tasks, err := taskService.Create(&input)

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, ProcessingError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, tasks)
		return
	}
}

// FindOneTask -
func FindOneTask(session db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		defer session.Close()

		taskService := NewTaskService(session)
		task, err := taskService.FindOne(vars["taskID"])

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, ProcessingError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, task)
		return
	}
}

// DisableTask -
func DisableTask(session db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		defer session.Close()

		taskService := NewTaskService(session)
		result, err := taskService.Disable(vars["taskID"])

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, ProcessingError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, result)
		return
	}
}

// DeleteTask -
func DeleteTask(session db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		defer session.Close()

		taskService := NewTaskService(session)
		err := taskService.Delete(vars["taskID"])

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, ProcessingError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, "done")
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

func respondWithError(w http.ResponseWriter, code int, errorType HTTPError, message string) {
	response, _ := json.Marshal(&errorHelper{Error: errorType.String(), Description: message})

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

func jsonDecode(input interface{}, reqBody io.ReadCloser) (*interface{}, error) {
	decoder := json.NewDecoder(reqBody)

	if err := decoder.Decode(&input); err != nil {
		return nil, err
	}

	return &input, nil
}
