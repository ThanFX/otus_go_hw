package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
	ErrWorkersFounded      = errors.New("no workers found")
)

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	ch := make(chan Task)
	var err error
	var eCount int32
	wg := sync.WaitGroup{}
	wg.Add(n)

	if n < 1 {
		return ErrWorkersFounded
	}

	for i := 0; i < n; i++ {
		go func() {
			for {
				v, ok := <-ch
				if !ok {
					break
				}
				err := v()
				if err != nil {
					atomic.AddInt32(&eCount, 1)
				}
			}
			wg.Done()
		}()
	}

	for _, t := range tasks {
		if m >= 0 && atomic.LoadInt32(&eCount) >= int32(m) {
			err = ErrErrorsLimitExceeded
			break
		}
		ch <- t
	}
	close(ch)

	wg.Wait()
	return err
}
