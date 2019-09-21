package tasker

import (
	"fmt"
	"log"
	"net/http"

	"github.com/roger-king/tasker/pkg"
)

// Tasker -
type Tasker struct {
	MongoConnection *pkg.MongoConnection
}

// Task - object
type Task struct {
	Name string `bson:"name"`
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

// Start -
func (c *Tasker) Start() {
	defer c.MongoConnection.DB.Close()
	log.Println("Starting tasker")
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

// HelloServer -
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
