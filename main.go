package main

import (
	"container/heap"
	"fmt"
)

// Task represents a task in the task management system
type Task struct {
	ID       int
	Priority int
	Name     string
}

// PriorityQueue implements a priority queue for tasks
type PriorityQueue []*Task

// Len returns the length of the priority queue
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// Less defines the ordering of tasks based on their priority
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority > pq[j].Priority // Higher priority tasks are placed before lower priority tasks
}

// Swap swaps the position of two tasks in the priority queue
func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].ID = i
	(*pq)[j].ID = j
}

// Push adds a task to the priority queue
func (pq *PriorityQueue) Push(x interface{}) {
	task := x.(*Task)
	task.ID = len(*pq)
	*pq = append(*pq, task)
}

// Pop removes and returns the highest priority task from the priority queue
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	task := old[n-1]
	task.ID = -1 // Mark the task as removed
	*pq = old[0 : n-1]
	return task
}

func main() {
	// Create a priority queue and add tasks
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	task1 := &Task{ID: 0, Priority: 3, Name: "Task A"}
	heap.Push(&pq, task1)

	task2 := &Task{ID: 1, Priority: 1, Name: "Task B"}
	heap.Push(&pq, task2)

	task3 := &Task{ID: 2, Priority: 2, Name: "Task C"}
	heap.Push(&pq, task3)

	// Process tasks in priority order
	for pq.Len() > 0 {
		task := heap.Pop(&pq).(*Task)
		fmt.Printf("Processing Task %d: %s\n", task.ID, task.Name)
	}
}
