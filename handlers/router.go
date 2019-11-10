package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/roger-king/tasker/services"
	"github.com/roger-king/tasker/utils"
)

func NewRouter(taskService *services.TaskService) *mux.Router {
	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/tasker").Subrouter()
	apiRouter.HandleFunc("/tasks", ListTasks(taskService)).Methods("GET")
	apiRouter.HandleFunc("/tasks", CreateTask(taskService)).Methods("POST")

	// Single Task Routes
	apiRouter.HandleFunc("/tasks/{taskID}", FindTask(taskService)).Methods("GET")
	apiRouter.HandleFunc("/tasks/{taskID}/disable", DisableTask(taskService)).Methods("PATCH")
	apiRouter.HandleFunc("/tasks/{taskID}", DeleteTask(taskService)).Methods("DELETE")

	// Web Admin - We have a reverse proxy for working on local developer :)
	r.PathPrefix("/static/").HandlerFunc(ServeWebAdmin)
	apiRouter.PathPrefix("/admin").HandlerFunc(ServeWebAdmin)

	return r
}

// Response Helpers:
type errorHelper struct {
	Error       string `json:"error"`
	Description string `json:"description"`
}

type dataHelper struct {
	Data interface{} `json:"data"`
}

func respondWithError(w http.ResponseWriter, code int, errorType utils.HTTPError, message string) {
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
