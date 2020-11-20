package workergroup

import (
	"sync"
	"testing"
	"time"
)

func Test_WorkGrouo(t *testing.T) {
	wg := NewWorkGroup(50)

	p := 0
	l := sync.Mutex{}

	for i := 0; i < 1000; i++ {
		wg.Add(func() {
			time.Sleep(time.Second / 10)
			l.Lock()
			p++
			l.Unlock()
		})
	}

	wg.Wait()

	if p != 1000 {
		t.Log("test failed; p != 1000", p)
		t.Fail()
	}
}
