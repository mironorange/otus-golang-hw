package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrWorkersLimited      = errors.New("workers are limited")
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
)

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var errorsCount int32
	if len(tasks) == 0 {
		return nil
	}
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}
	if n <= 0 {
		return ErrWorkersLimited
	}
	wg := sync.WaitGroup{}
	wg.Add(n)
	tasksChannel := make(chan Task, len(tasks))
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for task := range tasksChannel {
				if atomic.LoadInt32(&errorsCount) > int32(m) {
					break
				}
				if err := task(); err != nil {
					atomic.AddInt32(&errorsCount, 1)
				}
			}
		}()
	}
	for _, task := range tasks {
		tasksChannel <- task
	}
	close(tasksChannel)
	wg.Wait()
	if errorsCount > int32(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
