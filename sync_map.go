package go_sharded_mutex

import (
	"sync"
)

type SyncMap struct {
	counter *sync.Map
}

func newSyncMap(concurency int) *SyncMap {
	m := &sync.Map{}
	for i := 0; i < concurency; i++ {
		m.LoadOrStore(i, &Counter{
			mu: &sync.Mutex{},
		})
	}
	return &SyncMap{
		counter: m,
	}
}

func (a *SyncMap) incr(shard int) {
	s, _ := a.counter.Load(shard)
	s.(*Counter).safeIncr()
}
