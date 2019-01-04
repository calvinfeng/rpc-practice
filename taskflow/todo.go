package taskflow

import (
	"fmt"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

// NewTodo returns a Todo.
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
}

// AddTask adds a new task node to a todo list.
func (t *Todo) AddTask(n string) graph.Node {
	ctype := And
	name := n
	outcome := make(chan ExecutionResult)
	ready := make(chan int64)

	node := &Task{t.graph.NewNode(), ctype, name, ready, outcome}
	t.graph.AddNode(node)

	return node
}

// AddRelation creates dependency relationship between two tasks.
func (t *Todo) AddRelation(from, to graph.Node, et EdgeType) graph.Edge {
	etype := et

	edge := &Relation{t.graph.NewEdge(from, to), etype}
	t.graph.SetEdge(edge)

	return edge
}

// Traverse will walk through nodes.
func (t *Todo) Traverse() error {
	adj := t.adjacency()
	visited := make(map[int64]struct{})

	// Breadth first traversal
	queue := []int64{}
	for _, r := range t.Roots() {
		queue = append(queue, r.ID())
	}

	for len(queue) > 0 {
		id := queue[0]
		queue = queue[1:]

		task, ok := t.graph.Node(id).(*Task)
		if !ok {
			return fmt.Errorf("node %d is not a task", id)
		}

		if _, ok := visited[id]; ok {
			continue
		}

		visited[id] = struct{}{}
		fmt.Printf("working on task: %s\n", task.name)
		queue = append(queue, adj[id]...)
	}

	return nil
}

func (t *Todo) adjacency() map[int64][]int64 {
	adj := make(map[int64][]int64)

	nodeit := t.graph.Nodes()
	for nodeit.Next() {
		n := nodeit.Node()
		adj[n.ID()] = []int64{}
	}

	edgeit := t.graph.Edges()
	for edgeit.Next() {
		e := edgeit.Edge()
		adj[e.From().ID()] = append(adj[e.From().ID()], e.To().ID())
	}

	return adj
}

// Roots returns a list of node that has zero dependency.
func (t *Todo) Roots() []graph.Node {
	roots := []graph.Node{}

	nodeit := t.graph.Nodes()
	for nodeit.Next() {
		n := nodeit.Node()
		if t.graph.To(n.ID()).Len() == 0 {
			roots = append(roots, n)
		}
	}

	return roots
}

// Sort performs topological sorting.
func (t *Todo) Sort() error {
	nodes, err := topo.Sort(t.graph)
	if err != nil {
		return err
	}

	for _, n := range nodes {
		task, ok := n.(*Task)
		if !ok {
			return fmt.Errorf("node %d is not a task", n.ID())
		}

		fmt.Println(task.name)
	}

	return nil
}

// Cycles returns the number of cycles in Todo list.
func (t *Todo) Cycles() {
	fmt.Println("cycle count", len(topo.DirectedCyclesIn(t.graph)))
}
