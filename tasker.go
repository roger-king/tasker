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
	session := t.MongoConnection.DB
	log.Println("Starting tasker")

	// cr := cron.New()
	// cr.AddFunc(task.Schedule, func() {
	// 	fmt.Println("test")
	// })

	// cr.Start()

	r := mux.NewRouter()
	r.HandleFunc("/api/tasks", pkg.ListTasks(session)).Methods("GET")
	return r
}
