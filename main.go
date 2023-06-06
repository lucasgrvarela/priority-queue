package main

import (
	"container/heap"
	"fmt"
)

// Task represents a task in the task management system
type Task struct {
	ID          int
	Priority    int
	Description string
}

// PriorityQueue implements a priority queue for tasks
type PriorityQueue []*Task

// Len returns the length of the priority queue
func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push adds a task to the priority queue
func (pq *PriorityQueue) Push(task interface{}) {
	item := task.(*Task)
	*pq = append(*pq, item)
}

// Pop removes and returns the highest priority task from the priority queue
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func main() {
	// Create a priority queue and add tasks
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	task1 := &Task{ID: 0, Priority: 3, Description: "Do A"}
	heap.Push(&pq, task1)

	task2 := &Task{ID: 1, Priority: 1, Description: "Do B"}
	heap.Push(&pq, task2)

	task3 := &Task{ID: 2, Priority: 2, Description: "Do C"}
	heap.Push(&pq, task3)

	fmt.Println("Length of the priority queue is:", pq.Len())

	// Process tasks in highest priority order
	for pq.Len() > 0 {
		task := heap.Pop(&pq).(*Task)
		fmt.Printf("Processing task with Priority: %d, ID: %d, Description: '%s' \n", task.Priority, task.ID, task.Description)
	}
}
