package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/roger-king/tasker/services"
	"github.com/roger-king/tasker/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(taskService *services.TaskService, github *services.GithubAuthService, db *mongo.Client) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/check", CheckSession()).Methods("GET")

	apiRouter := r.PathPrefix("/tasker").Subrouter()
	apiRouter.Use(authMiddleware)
	apiRouter.HandleFunc("/tasks", ListTasks(taskService)).Methods("GET")
	apiRouter.HandleFunc("/tasks", CreateTask(taskService)).Methods("POST")

	// Single Task Routes
	apiRouter.HandleFunc("/tasks/{taskID}", FindTask(taskService)).Methods("GET")
	apiRouter.HandleFunc("/tasks/{taskID}/disable", DisableTask(taskService)).Methods("PATCH")
	apiRouter.HandleFunc("/tasks/{taskID}", DeleteTask(taskService)).Methods("DELETE")

	// Web Admin - We have a reverse proxy for working on local developer :)
	r.PathPrefix("/static/").HandlerFunc(ServeWebAdmin)
	r.PathPrefix("/images/").HandlerFunc(ServeWebAdmin)
	r.PathPrefix("/tasker/admin").HandlerFunc(ServeWebAdmin)

	// authenticate route
	// Public routes for github access
	oauth := r.PathPrefix("/oauth").Subrouter()
	oauth.HandleFunc("/authenticate/{code}", LoginHandler(github, db)).Methods("POST")
	oauth.HandleFunc("/github/user", FetchUserClientIDHandler(github)).Methods("GET")

	return r
}

// Response Helpers:
type errorHelper struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

type dataHelper struct {
	Data interface{} `json:"data"`
}

func respondWithError(w http.ResponseWriter, code int, errorType utils.HTTPError, message interface{}) {
	response, _ := json.Marshal(&errorHelper{Error: errorType.String(), Data: message})

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
