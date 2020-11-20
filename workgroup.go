package workergroup

import (
	"sync"
)

// WorkGroup runs a given number of functions concurrently.
type WorkGroup struct {
	main  chan func()
	waitg *sync.WaitGroup
}

// NewWorkGroup allocates and returns a new WorkGroup.
func NewWorkGroup(n int) *WorkGroup {

	wg := &WorkGroup{
		main:  make(chan func()),
		waitg: &sync.WaitGroup{},
	}

	wg.waitg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.waitg.Done()
			for f := range wg.main {
				f()
			}
		}()
	}

	return wg

}

// Add adds a new worker function to execute.
func (wg *WorkGroup) Add(f func()) {
	wg.main <- f
}

// Wait waits until all worker function are finished.
func (wg *WorkGroup) Wait() {
	close(wg.main)
	wg.waitg.Wait()
}
