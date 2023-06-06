# Priority Queue in Go

This Go program demonstrates the implementation of a priority queue using the package [container/heap](https://pkg.go.dev/container/heap). The program initializes a priority queue and adds several tasks with different priorities. The tasks are then processed in priority order, with the highest priority task being processed first.

## Usage

1. Make sure you have Go installed on your system.

2. Clone the repository:

```bash
$ git clone https://github.com/lucasgrvarela/priority-queue.git
```

3. Navigate to the project directory:
```bash
$ cd priority-queue
```

4. Run the program:
```bash
$ go run main.go
Length of the priority queue is: 3
Processing task with Priority: 3, ID: 0, Description: 'Do A'
Processing task with Priority: 2, ID: 2, Description: 'Do C'
Processing task with Priority: 1, ID: 1, Description: 'Do B'
Length of the priority queue is: 0
```

Note: The methods Len(), Less(), Swap(), Push() and Pop() are implemented to satisfy the [heap.Interface](https://pkg.go.dev/container/heap#Interface)
