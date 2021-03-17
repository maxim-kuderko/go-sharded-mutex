package go_sharded_mutex

import (
	"sync"
)

type Counter struct {
	Count int64
	mu    *sync.Mutex
}

func (c *Counter) safeIncr() {
	c.mu.Lock()
	c.incr()
	c.mu.Unlock()
}

func (c *Counter) incr() {
	c.Count++
}
