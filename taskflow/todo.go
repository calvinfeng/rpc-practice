package taskflow

import (
	"fmt"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/traverse"
)

func NewTodo(n string) *Todo {
	return &Todo{
		graph: simple.NewDirectedGraph(),
		name:  n,
	}
}

// Todo is a DAG.
type Todo struct {
	graph *simple.DirectedGraph
	name  string
	root  graph.Node
}

// Root returns the root node of a task DAG, which is the node that has no from edges.
func (t *Todo) Root() graph.Node {
	return t.root
}

func (t *Todo) AddTask(n string) graph.Node {
	node := t.graph.NewNode()

	if t.root == nil {
		t.root = node
	}

	t.graph.AddNode(node)

	return &Task{id: node.ID(), name: n}
}

func (t *Todo) AddRelation(from, to graph.Node) graph.Edge {
	return &Relation{simple.Edge{F: from, T: to}}
}

func (t *Todo) Traverse() {
	bfs := traverse.BreadthFirst{}
	for n := bfs.Walk(t.graph, t.root, func(n graph.Node, d int) bool {
		return true
	}); n != nil; {
		fmt.Println(n)
	}
}
