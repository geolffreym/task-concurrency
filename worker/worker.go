package worker

import (
	"concurrent/task"
	"fmt"
	"sync/atomic"
	"time"
)

type Worker[T any] struct {
	Tasks    []*task.Task[T]
	Channels chan *task.Result
	Rate     int
	pool     int32
}

// Run go routines and populate channels based on defined tasks
func (w *Worker[T]) Run() {
	// Buffered channels
	var tasks []*task.Task[T] = w.Tasks
	// defer close(channels)

	//Make thread callk
	for i, currentTask := range tasks {
		fmt.Printf("\nRunning task %d from %d", i+1, len(tasks))
		// Sync channels
		atomic.AddInt32(&w.pool, 1)
		// How long to wait?
		time.Sleep(time.Duration(w.Rate))

		go func(index int, task_ *task.Task[T]) {
			w.Channels <- &task.Result{
				Thread:  index + 1,
				Payload: task_.Call(),
			}
		}(i, currentTask)
	}
}

// Consume channels from Worker 
func (w *Worker[T]) Process() {
	//For each channel response
	for {
		select {
		case response := <-w.Channels:
			fmt.Printf("\nTasks #%d with payload %d", response.Thread, response.Payload)
			atomic.AddInt32(&w.pool, -1)

		default:
			// Stop when decreasing threads == 0
			if atomic.LoadInt32(&w.pool) == 0 {
				return
			}

		}
	}
}
