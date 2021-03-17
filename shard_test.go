package go_sharded_mutex

import (
	"runtime"
	"sync"
	"testing"
)

func Benchmark_Array(b *testing.B) {

	concurrency := runtime.GOMAXPROCS(0)
	benchmarks := []struct {
		name string
		fn   func(shard int)
	}{
		{"Atomic", newAtomic().incr},
		{"ShardedAtomic", newShardedAtomic(concurrency).incr},
		{"Array", newArray(concurrency).incr},
		{"SyncMap", newSyncMap(concurrency).incr},
		{"SyncPool", newSyncPool().incr},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			wg := sync.WaitGroup{}
			wg.Add(concurrency * 4)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < concurrency*4; i++ {
				go func(i int) {
					defer wg.Done()
					for j := 0; j < b.N; j++ {
						bm.fn(i % concurrency)
					}
				}(i)
			}
			wg.Wait()
		})
	}
}
