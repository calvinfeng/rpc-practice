package taskflow

import (
	"fmt"
	"time"

	"gonum.org/v1/gonum/graph"
)

type (
	// EdgeType defines outgoing edge condition.
	EdgeType int

	// ConnectorType describes the condition of how a task should aggregate all incoming edges.
	ConnectorType int

	// ExecutionResult captures the outcome of an execution on a task.
	ExecutionResult struct {
		msg string
		err error
		id  int64
	}

	// Task satisfies simple.Node interface.
	Task struct {
		graph.Node
		ctype   ConnectorType
		name    string
		ready   chan int64
		outcome chan ExecutionResult
	}

	// Relation satisfies simple.Edge interface.
	Relation struct {
		graph.Edge
		etype EdgeType
	}
)

func (t *Task) activate(upstreams []<-chan ExecutionResult) {
	out := fanIn(upstreams...)

	switch t.ctype {
	case And:
		for i := 0; i < len(upstreams); i++ {
			<-out
		}

		t.ready <- t.ID()
	case Or:
		for i := 0; i < len(upstreams); i++ {
			<-out
			t.ready <- t.ID()
			break
		}
	}
}

func fanIn(inputs ...<-chan ExecutionResult) <-chan ExecutionResult {
	out := make(chan ExecutionResult)

	for _, in := range inputs {
		go func(ch <-chan ExecutionResult) {
			out <- <-ch
		}(in)
	}

	return out
}

func (t *Task) execute() {
	fmt.Println("starting", t.name)

	time.Sleep(1 * time.Second)

	res := ExecutionResult{id: t.ID()}
	// if rand.Float64() > 0.90 {
	// 	res.err = errors.New("encountered a random error")
	// } else {
	// 	res.msg = fmt.Sprintf("completed %s", t.name)
	// }
	res.msg = fmt.Sprintf("completed %s", t.name)

	fmt.Println("done with", t.name)

	t.outcome <- res
}

// List of EdgeType.
const (
	Success = iota
	Error
	Done
)

// List of ConnectorType
const (
	And ConnectorType = iota
	Or
)

func (et EdgeType) String() string {
	var names = map[EdgeType]string{
		Success: "Success",
		Error:   "Error",
		Done:    "Done",
	}

	return names[et]
}
