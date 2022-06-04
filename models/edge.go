package models

import (
	"fmt"
	"sort"
	"strings"
)

//Edge type
type Edge struct {
	Key     string
	Vertexs []string
}

// NewEdge inits a new Edge instance
func NewEdge(vertexs ...string) Edge {
	sort.Strings(vertexs)
	key := strings.Join(vertexs, "_")
	return Edge{key, vertexs}
}

// String representation of a Link
func (e Edge) String() string {
	return fmt.Sprintf("key=%s nodes=%s\n", e.Key, e.Vertexs)
}
