package main

import (
	"concurrent/task"
	"concurrent/worker"
	"sort"
)

//
func main() {
	// Create a task to processing
	// TODO run as microservices?
	// Add queue features
	sum := &task.Task[int]{
		Input: []int{1, 2, 3},
		Implementation: func(num ...int) int {
			var sum int = 0
			for _, n := range num {
				sum += n
			}

			return sum
		},
	}

	sub := &task.Task[int]{
		Input: []int{10, 9, 8, 8},
		Implementation: func(num ...int) int {
			sort.Ints(num)                  // Sort Input then get bigger number
			var minus int = num[len(num)-1] // Get bigger number
			// Subtract from bigger number
			for _, n := range num[:len(num)-1] {
				minus -= n
			}

			return minus
		},
	}

	// Run concurrent tasks using channels and routines
	tasks := []*task.Task[int]{sum, sub}
	channels := make(chan *task.Result, len(tasks))

	worker := worker.Worker[int]{Tasks: tasks, Channels: channels}
	worker.Run()     // Run worker
	worker.Process() // Process channel response
}
