package go_sharded_mutex

import (
	"go.uber.org/atomic"
)

type ShardedAtomic struct {
	counter []*atomic.Int64
}

func newShardedAtomic(c int) *ShardedAtomic {
	atomics := make([]*atomic.Int64, 0, c)
	for i := 0; i < c; i++ {
		atomics = append(atomics, atomic.NewInt64(0))
	}
	return &ShardedAtomic{
		counter: atomics,
	}
}

func (a *ShardedAtomic) incr(shard int) {
	a.counter[shard].Inc()
}
