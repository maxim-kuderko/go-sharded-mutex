package go_sharded_mutex

import (
	"sync"
)

type Array struct {
	counters []*Counter

	size uint32
}

func newArray(c int) *Array {
	counters := make([]*Counter, 0, c)
	mutexes := make([]*sync.Mutex, 0, c)
	for i := 0; i < c; i++ {
		counters = append(counters, &Counter{mu: &sync.Mutex{}})
		mutexes = append(mutexes, &sync.Mutex{})
	}
	return &Array{
		counters: counters,
		size:     uint32(c),
	}
}

func (a *Array) incr(shard int) {
	a.counters[shard].safeIncr()
}

func (a *Array) get() int64 {
	output := int64(0)
	for _, i := range a.counters {
		output += i.Count
	}
	return output
}
