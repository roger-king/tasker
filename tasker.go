package tasker

import (
	"log"

	"github.com/gorilla/mux"
	cron "github.com/robfig/cron/v3"
	"github.com/roger-king/tasker/pkg"
	db "upper.io/db.v3"
)

// Tasker -
type Tasker struct {
	DB        db.Database
	Scheduler *cron.Cron
}

type TaskerConfig struct {
	Details *pkg.ConnectionDetails `required:"true"`
}

// New - Creates a new instance of tasker
func New(tc *TaskerConfig) *Tasker {
	m, err := pkg.NewMongoConnection()

	if err != nil {
		log.Panic(err)
	}

	return &Tasker{
		Scheduler: cron.New(),
		DB:        m,
	}
}

// Start - returns a mux router instance
func (t *Tasker) Start() *mux.Router {
	t.Scheduler.Start()

	taskService := &pkg.TaskService{
		DB:        t.DB,
		Scheduler: t.Scheduler,
	}

	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/tasks", pkg.ListTasks(taskService)).Methods("GET")
	apiRouter.HandleFunc("/tasks", pkg.CreateTask(taskService)).Methods("POST")

	// Single Task Routes
	apiRouter.HandleFunc("/tasks/{taskID}", pkg.FindTask(taskService)).Methods("GET")
	// apiRouter.HandleFunc("/tasks/{taskID}/disable", pkg.DisableTask(session)).Methods("PATCH")
	apiRouter.HandleFunc("/tasks/{taskID}", pkg.DeleteTask(taskService)).Methods("DELETE")
	return r
}
