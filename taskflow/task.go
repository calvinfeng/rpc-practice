package taskflow

import "gonum.org/v1/gonum/graph"

// Task is satisfies simple.Node interface
type Task struct {
	graph.Node
	name string
}

// Relation satisfies simple.Edge interface
type Relation struct {
	graph.Edge
	strength float64
}
