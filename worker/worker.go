package worker

import (
	"concurrent/task"
	"fmt"
)

type Worker[T any] struct {
	Tasks []*task.Task[T]
}

// Run go routines and populate channels based on defined tasks
func (w Worker[T]) Run() chan *task.Result {
	// Buffered channels
	var tasks []*task.Task[T] = w.Tasks
	var channels chan *task.Result = make(chan *task.Result, len(w.Tasks))

	//Make thread call
	for i, currentTask := range tasks {
		fmt.Printf("\nRunning task %d from %d", i+1, len(tasks))
		go func(index int, task_ *task.Task[T], ch chan *task.Result) {
			ch <- &task.Result{
				Thread:   index + 1,
				Channels: len(tasks),
				Payload:  task_.Call(),
			}
		}(i, currentTask, channels)
	}

	return channels
}

// Consume channels from Worker and return results
func (w Worker[T]) Process(ch chan *task.Result) []*task.Result {
	//Array of responses coming from channel
	responses := make([]*task.Result, 0)

	//For each channel response
	for response := range ch {
		fmt.Printf("\nTasks #%d with payload %d", response.Thread, response.Payload)
		responses = append(responses, response)
		if len(responses) == response.Channels {
			close(ch)
		}
	}

	return responses

}
