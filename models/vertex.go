package models

import "fmt"

type Vertex struct {
	Name          string
	SimConditions map[string]bool
	Edges         []*Edge
	Vertexs       map[string]*Vertex
}

func NewVertex(name string) Vertex {
	return Vertex{
		Name:          name,
		SimConditions: make(map[string]bool),
		Edges:         make([]*Edge, 0),
		Vertexs:       make(map[string]*Vertex),
	}
}

func (v *Vertex) Connect(otherVertex *Vertex) *Edge {
	edge := NewEdge(v.Name, otherVertex.Name)
	return v.ConnectVia(&edge, otherVertex)
}

// ConnectVia via link to Node
func (v *Vertex) ConnectVia(edge *Edge, other *Vertex) *Edge {
	if v.Vertexs[edge.Key] == nil {
		v.Edges = append(v.Edges, edge)
		v.Vertexs[edge.Key] = other
	}
	return edge
}

// String representation for a Node
func (v *Vertex) String() string {
	var links string
	for k, v := range v.Vertexs {
		links += fmt.Sprintf("%s:%s ", k, v.Name)
	}
	if len(v.Vertexs) > 0 {
		links = links[:len(links)-1]
	}
	return fmt.Sprintf("name=%s links=map[%s]", v.Name, links)
}
