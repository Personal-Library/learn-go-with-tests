package sync

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

// Any goroutine calling Inc will acquire the lock on Counter if they are first.
// All other goroutines will have to wait for it to be unlocked before going
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
