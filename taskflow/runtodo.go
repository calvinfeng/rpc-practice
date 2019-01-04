package taskflow

import "fmt"

// RunTodo executes the todo list.
func RunTodo(t *Todo) {
	adj := t.adjacency()

	queue := NewActiveQueue(t.graph)
	for _, node := range t.Roots() {
		task, ok := node.(*Task)
		if !ok {
			panic(fmt.Sprintf("node %d is not a task!", node.ID()))
		}
		queue.add(task)
	}

	for queue.notEmpty() {
		task := queue.next()
		fmt.Println("current task popped from queue", task.name)

		go task.execute()

		for _, childID := range adj[task.ID()] {
			node := t.graph.Node(childID)
			if node == nil {
				panic(fmt.Sprintf("node %d does not exist", childID))
			}

			child, ok := node.(*Task)
			if !ok {
				panic(fmt.Sprintf("node %d is not a task", node.ID()))
			}

			queue.add(child)
		}
	}
}
