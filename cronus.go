package cronus

import (
	"fmt"
	"log"

	"github.com/roger-king/cronus/pkg"
)

// Cronus -
type Cronus struct {
	MongoConnection *pkg.MongoConnection
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
	fmt.Println("Starting Cronus Server.")
}
