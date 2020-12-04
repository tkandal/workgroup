package workergroup

import (
	"sync"
)

// WorkGroup runs a given number of functions concurrently.
type WorkGroup struct {
	workQueue  chan func()
	waitg *sync.WaitGroup
}

// NewWorkGroup allocates and returns a new WorkGroup.
func NewWorkGroup(n int) *WorkGroup {
	wg := &WorkGroup{
		workQueue:  make(chan func(), n),
		waitg: &sync.WaitGroup{},
	}

	for i := 0; i < n; i++ {
		go func() {
			for f := range wg.workQueue {
				f()
				wg.waitg.Done()
			}
		}()
	}
	return wg
}

// Add adds a new worker function to execute.
func (wg *WorkGroup) Add(f func()) {
	wg.waitg.Add(1)
	wg.workQueue <- f
}

// Wait waits until all worker function are finished.
func (wg *WorkGroup) Wait() {
	wg.waitg.Wait()
	close(wg.workQueue)
}
