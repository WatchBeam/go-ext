package msync

import (
	"sync"
	"sync/atomic"
)

// Parallel spawns `workers` goroutines and invokes the function in each
// routine from `start` to `end`. It blocks until all functions return.
func Parallel(count, workers int, fn func(i int)) {
	ParallelChunks(count, 1, workers, func(start, _ int) { fn(start) })
}

// ParallelRanges divides up the the domain into segments and dispatches
// them to worker functions.
func ParallelRanges(count, workers int, fn func(start, end int)) {
	if count == 0 {
		return
	}
	if count < workers {
		fn(0, count)
		return
	}

	segment := count / workers

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		a := i * segment
		b := (i + 1) * segment
		if i == workers-1 {
			b = count
		}
		if b-a == 0 {
			continue
		}

		wg.Add(1)
		go func() {
			fn(a, b)
			wg.Done()
		}()
	}

	wg.Wait()
}

// ParallelChunks divides the domain into chunks of "chunkSize" and feeds it
// to the function, with bounded concurrency of `worker` functions.
func ParallelChunks(count, chunkSize, workers int, fn func(start, end int)) {
	if count == 0 {
		return
	}
	if chunkSize == 0 {
		panic("Cannot pass a chunkSize of 0 to ParallelChunks")
	}
	if count < chunkSize {
		fn(0, count)
		return
	}

	var (
		wg      sync.WaitGroup
		offset  int64
		chunk64 = int64(chunkSize)
	)

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			for {
				upper := int(atomic.AddInt64(&offset, chunk64))
				lower := upper - chunkSize
				if lower >= count {
					wg.Done()
					return
				}
				if upper > count {
					upper = count
				}

				fn(lower, upper)
			}
		}()
	}

	wg.Wait()

}
