package cronus

import (
	"fmt"
)

// Cronus -
type Cronus struct {
	Name string
}

// New -
func New(name string) *Cronus {
	return &Cronus{
		Name: name,
	}
}

// Say -
func (c *Cronus) Say() {
	fmt.Println("Hello", c.Name)
}
