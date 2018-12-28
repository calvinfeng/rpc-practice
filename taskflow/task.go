package taskflow

import "gonum.org/v1/gonum/graph"

// Task is satisfies simple.Node interface
type Task struct {
	id   int64
	name string
}

func (t *Task) ID() int64 {
	return t.id
}

// Relation satisfies simple.Edge interface
type Relation struct {
	graph.Edge
}
