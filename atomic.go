package go_sharded_mutex

import (
	"go.uber.org/atomic"
)

type Atomic struct {
	counter *atomic.Int64
}

func newAtomic() *Atomic {
	return &Atomic{
		counter: atomic.NewInt64(0),
	}
}

func (a *Atomic) incr(shard int) {
	a.counter.Inc()
}
