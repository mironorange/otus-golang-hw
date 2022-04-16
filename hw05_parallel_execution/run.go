package hw05parallelexecution

import (
	"errors"
	"sync/atomic"
)

var (
	ErrWorkersLimited      = errors.New("workers are limited")
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
)

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var errorsCount, runTasksCount int32
	if len(tasks) == 0 {
		return nil
	}
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}
	if n <= 0 {
		return ErrWorkersLimited
	}

	tasksChannel := make(chan Task, len(tasks))
	errorsChannel := make(chan struct{}, n)
	terminationChannel := make(chan struct{}, n)

	for i := 0; i < n; i++ {
		go func() {
			// Если не было ошибок, то горутина завершит свою работу, так как закроется канал.
			for task := range tasksChannel {
				if err := task(); err != nil {
					errorsChannel <- struct{}{}
				}
				terminationChannel <- struct{}{}
				// Если количество ошибок больше ожидаемых, то завершить чтение из канала.
				// И завершить работу горутины
				if errorsCount > int32(m) {
					break
				}
			}
		}()
	}

	for _, task := range tasks {
		tasksChannel <- task
	}
	close(tasksChannel)

	for {
		// Если не добавлена секция default, то оператор select ждет сообщение от одного из канала.
		select {
		// Считаем количество завершившихся задач.
		case <-terminationChannel:
			atomic.AddInt32(&runTasksCount, 1)
		// Считаем количество возникших ошибок.
		case <-errorsChannel:
			atomic.AddInt32(&errorsCount, 1)
		// Контролируем количество завершившихся задач и количество ошибок.
		default:
			if runTasksCount == int32(len(tasks)) {
				return nil
			}
			if errorsCount > int32(m) {
				return ErrErrorsLimitExceeded
			}
		}
	}
}
