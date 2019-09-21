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
func New() *Cronus {
	m, err := pkg.NewMongoConnection()

	if err != nil {
		log.Panic(err)
	}

	return &Cronus{
		MongoConnection: m,
	}
}

// Start -
func (c *Cronus) Start() {
	defer c.MongoConnection.DB.Close()
	log.Println("Starting cronus")
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
