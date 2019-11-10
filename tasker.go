package tasker

import (
	"os"

	"github.com/gorilla/mux"
	cron "github.com/robfig/cron/v3"
	"github.com/roger-king/tasker/pkg"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

// Tasker -
type Tasker struct {
	DB        *mongo.Client
	Scheduler *cron.Cron
}

type TaskerConfig struct {
	MongoConnectionURL string `required:"true"`
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

// New - Creates a new instance of tasker
func New(tc *TaskerConfig) *Tasker {
	m, err := pkg.NewMongoConnection(tc.MongoConnectionURL)

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
	log.Info("Starting Tasker application")
	t.Scheduler.Start()

	taskService := &pkg.TaskService{
		DB:        t.DB,
		Scheduler: t.Scheduler,
	}

	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/tasker").Subrouter()
	apiRouter.HandleFunc("/tasks", pkg.ListTasks(taskService)).Methods("GET")
	apiRouter.HandleFunc("/tasks", pkg.CreateTask(taskService)).Methods("POST")

	// Single Task Routes
	apiRouter.HandleFunc("/tasks/{taskID}", pkg.FindTask(taskService)).Methods("GET")
	apiRouter.HandleFunc("/tasks/{taskID}/disable", pkg.DisableTask(taskService)).Methods("PATCH")
	apiRouter.HandleFunc("/tasks/{taskID}", pkg.DeleteTask(taskService)).Methods("DELETE")

	// Web Admin - We have a reverse proxy for working on local developer :)
	r.PathPrefix("/static/").HandlerFunc(pkg.ServeWebAdmin)
	apiRouter.PathPrefix("/admin").HandlerFunc(pkg.ServeWebAdmin)
	return r
}
