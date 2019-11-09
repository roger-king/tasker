package tasker

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/robfig/cron"
	"github.com/roger-king/tasker/pkg"
	db "upper.io/db.v3"
)

// Tasker -
type Tasker struct {
	Client    *redis.Client
	DB        db.Database
	Scheduler *cron.Cron
}

type TaskerConfig struct {
	ConnectionType pkg.ConnectionType     `required:"true"`
	Details        *pkg.ConnectionDetails `required:"true"`
}

// New - Creates a new instance of tasker
func New(tc *TaskerConfig) *Tasker {
	var client *redis.Client

	switch tc.ConnectionType {
	// case pkg.MONGO:
	// 	m, err := pkg.NewMongoConnection()

	// 	if err != nil {
	// 		log.Panic(err)
	// 	}
	// 	t.TaskService.DB = m
	// 	break
	case pkg.REDIS:
		r, err := pkg.NewRedisConnection(tc.Details)

		if err != nil {
			log.Panic(err)
		}
		client = r
		break
	default:
		// The default connection type will be redis
		log.Panic("Please provide a valid connection type")
		break
	}

	return &Tasker{
		Scheduler: cron.New(),
		Client:    client,
	}
}

// Start - returns a mux router instance
func (t *Tasker) Start() *mux.Router {
	log.Println("Starting tasker")
	t.Scheduler.Start()
	taskService := &pkg.TaskService{
		Client:    t.Client,
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
