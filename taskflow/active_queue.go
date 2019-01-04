package taskflow

import (
	"fmt"

	"gonum.org/v1/gonum/graph/simple"
)

// NewActiveQueue returns an active queue.
func NewActiveQueue(dg *simple.DirectedGraph) *ActiveQueue {
	return &ActiveQueue{
		graph:    dg,
		nextNode: make(chan int64),
		set:      make(map[int64]*Task),
	}
}

// ActiveQueue implements first-ready-first-out policy. A node is ready when it has all its
// dependencies met.
type ActiveQueue struct {
	graph    *simple.DirectedGraph
	nextNode chan int64
	set      map[int64]*Task
}

func (q *ActiveQueue) next() *Task {
	id := <-q.nextNode

	n, ok := q.set[id]
	if !ok {
		panic("node is not found in queue")
	}

	delete(q.set, id)
	return n
}

func (q *ActiveQueue) has(t *Task) bool {
	_, ok := q.set[t.ID()]
	return ok
}

func (q *ActiveQueue) add(t *Task) {
	if _, ok := q.set[t.ID()]; ok {
		return
	}

	q.set[t.ID()] = t

	upstreams := []<-chan ExecutionResult{}

	iter := q.graph.To(t.ID())
	for iter.Next() {
		node := iter.Node()
		parent, ok := node.(*Task)
		if !ok {
			panic("parent is not a task!")
		}

		edge := q.graph.Edge(parent.ID(), t.ID())
		relation, ok := edge.(*Relation)
		if !ok {
			panic("edge is not a relation!")
		}

		switch relation.etype {
		case Success:
			upstreams = append(upstreams, filterSuccessUpstream(parent.outcome))
		case Error:
			upstreams = append(upstreams, filterErrorUpstream(parent.outcome))
		case Done:
			upstreams = append(upstreams, filterDoneUpstream(parent.outcome))
		}
	}

	go t.activate(upstreams)

	go func(next chan<- int64, ready <-chan int64) {
		r := <-ready
		fmt.Printf("%d is ready\n", r)
		next <- r
	}(q.nextNode, t.ready)
}

func filterSuccessUpstream(outcome <-chan ExecutionResult) <-chan ExecutionResult {
	ch := make(chan ExecutionResult)
	go func(read <-chan ExecutionResult, write chan<- ExecutionResult) {
		res := <-read
		if res.err == nil {
			write <- res
		}
	}(outcome, ch)

	return ch
}

func filterErrorUpstream(outcome <-chan ExecutionResult) <-chan ExecutionResult {
	ch := make(chan ExecutionResult)
	go func(read <-chan ExecutionResult, write chan<- ExecutionResult) {
		res := <-read
		if res.err != nil {
			write <- res
		}
	}(outcome, ch)

	return ch
}

func filterDoneUpstream(outcome <-chan ExecutionResult) <-chan ExecutionResult {
	ch := make(chan ExecutionResult)
	go func(read <-chan ExecutionResult, write chan<- ExecutionResult) {
		write <- <-read
	}(outcome, ch)

	return ch
}

func (q *ActiveQueue) notEmpty() bool {
	return len(q.set) != 0
}
