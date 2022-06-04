package models

import (
	"fmt"
	"sagaAlienInvasion/models"
)

//City struct contains model vertex with composition
//And contains map of road names and edges
type City struct {
	models.Vertex
	RoadNames map[string]string
}

// NewCity returns a new city struct with taken name
func NewCity(name string) City {
	return City{
		Vertex:    models.NewVertex(name),
		RoadNames: make(map[string]string),
	}
}

//IsDestroyed checks if City is destroyed
func (c *City) IsDestroyed() bool {
	return c.SimConditions[DestroyedCond]
}

// Destroy  makes City destroyed with change in SimConditions to true
func (c *City) Destroy() {
	c.SimConditions[DestroyedCond] = true
}

// String representation for a City does not print destroyed linked Cities
func (c *City) String() string {
	var edges string
	for _, edge := range c.Edges {
		n := c.Vertexs[edge.Key]
		other := City{Vertex: *n}
		// If other City destroyed print nothing
		if other.IsDestroyed() {
			continue
		}
		// If other City survived print Link
		edges += fmt.Sprintf("%s=%s ", c.RoadNames[edge.Key], other.Name)
	}
	if len(edges) == 0 {
		return c.Name
	}
	return fmt.Sprintf("%s %s", c.Name, edges[:len(edges)-1])
}
