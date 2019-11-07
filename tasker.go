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
	mongoConnection db.Database
	redisConnection *redis.Client
	Scheduler       *cron.Cron
}

type TaskerCongfig struct {
	ConnectionType pkg.ConnectionType
	Details        *pkg.ConnectionDetails
}

// New - Creates a new instance of tasker
func New(tc *TaskerCongfig) *Tasker {
	t := &Tasker{
		Scheduler: cron.New(),
	}

	switch tc.ConnectionType {
	case pkg.MONGO:
		m, err := pkg.NewMongoConnection()

		if err != nil {
			log.Panic(err)
		}
		t.mongoConnection = m
		break
	case pkg.REDIS:
		r, err := pkg.NewRedisConnection(tc.Details)

		if err != nil {
			log.Panic(err)
		}
		t.redisConnection = r
		break
	default:
		// The default connection type will be redis
		break
	}

	return t
}

// Start - returns a mux router instance
func (t *Tasker) Start() *mux.Router {
	log.Println("Starting tasker")
	t.Scheduler.Start()
	session := t.mongoConnection

	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/tasks", pkg.ListTasks(session)).Methods("GET")
	apiRouter.HandleFunc("/tasks", pkg.CreateTask(session, t.Scheduler)).Methods("POST")

	// test
	// Single Task Routes
	apiRouter.HandleFunc("/tasks/{taskID}", pkg.FindOneTask(session)).Methods("GET")
	apiRouter.HandleFunc("/tasks/{taskID}/disable", pkg.DisableTask(session)).Methods("PATCH")
	apiRouter.HandleFunc("/tasks/{taskID}", pkg.DeleteTask(session)).Methods("DELETE")
	return r
}
