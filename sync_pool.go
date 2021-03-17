package go_sharded_mutex

import (
	"sync"
)

type SyncPool struct {
	counter *sync.Pool
}

func newSyncPool() *SyncPool {
	return &SyncPool{
		counter: &sync.Pool{New: func() interface{} {
			return &Counter{
				mu: &sync.Mutex{},
			}
		}},
	}
}

func (a *SyncPool) incr(shard int) {
	a.counter.Get().(*Counter).safeIncr()
}
