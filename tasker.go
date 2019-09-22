package tasker

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/roger-king/tasker/pkg"
)

// Tasker -
type Tasker struct {
	MongoConnection *pkg.MongoConnection
}

// New -
func New() *Tasker {
	m, err := pkg.NewMongoConnection()

	if err != nil {
		log.Panic(err)
	}

	return &Tasker{
		MongoConnection: m,
	}
}

// Start - returns a mux router instance
func (t *Tasker) Start() *mux.Router {
	log.Println("Starting tasker")

	session := t.MongoConnection.DB

	// cr := cron.New()
	// cr.AddFunc(task.Schedule, func() {
	// 	fmt.Println("test")
	// })

	// cr.Start()

	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/tasks", pkg.ListTasks(session)).Methods("GET")
	apiRouter.HandleFunc("/tasks", pkg.CreateTask(session)).Methods("POST")
	apiRouter.HandleFunc("/tasks/{taskID}", pkg.DisableTask(session)).Methods("PATCH")
	return r
}
