package tasker

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/robfig/cron"
	"github.com/roger-king/tasker/pkg"
	db "upper.io/db.v3"
)

// Tasker -
type Tasker struct {
	mongoConnection db.Database
	Scheduler       *cron.Cron
}

// New -
func New() *Tasker {
	m, err := pkg.NewMongoConnection()

	if err != nil {
		log.Panic(err)
	}

	return &Tasker{
		mongoConnection: m,
		Scheduler:       cron.New(),
	}
}

// Start - returns a mux router instance
func (t *Tasker) Start() *mux.Router {
	log.Println("Starting tasker")
	t.Scheduler.Start()
	session := t.mongoConnection

	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/tasks", pkg.ListTasks(session)).Methods("GET")
	apiRouter.HandleFunc("/tasks", pkg.CreateTask(session)).Methods("POST")
	apiRouter.HandleFunc("/tasks/{taskID}", pkg.DisableTask(session)).Methods("PATCH")
	return r
}
