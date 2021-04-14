package counter

import "sync"

// Go types that are bound to the UI must be thread-safe, because each binding
// is executed in its own goroutine. In this simple case we may use atomic
// operations, but for more complex cases one should use proper synchronization.
type Counter struct {
	sync.Mutex
	count int
}

func (c *Counter) Add(n int) {
	c.Lock()
	defer c.Unlock()
	c.count = c.count + n
}

func (c *Counter) Value() int {
	c.Lock()
	defer c.Unlock()
	return c.count
}
