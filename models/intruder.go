package models

import "fmt"

// Intruder is a struct for storing intruder data
// TODO: add comments
type Intruder struct {
	Name          string
	SimConditions map[string]bool
	Vertex        *Vertex
}

// NewIntruder creates a new intruder
func NewIntruder(name string) Intruder {
	return Intruder{
		Name:          name,
		SimConditions: make(map[string]bool),
	}
}

// String returns a string representation of the intruder
func (i *Intruder) String() string {
	return fmt.Sprintf("name=%s node={%s}\n", i.Name, i.Vertex)
}
