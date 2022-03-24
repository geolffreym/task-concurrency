package task

type Result struct {
	Thread  int
	Payload any
}

type Task[T any] struct {
	Input          []T
	Implementation func(...T) T // The definition for task
}

func (task Task[T]) Call() T {
	return task.Implementation(task.Input...)
}
