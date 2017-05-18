package msync

import (
	"sync"
	"sync/atomic"
)

// Parallel spawns `workers` goroutines and invokes the function in each
// routine from `start` to `end`. It blocks until all functions return.
func Parallel(start, end, workers int, fn func(i int)) {
	if end-start == 0 {
		return
	}

	var wg sync.WaitGroup
	wg.Add(workers)

	current := uint64(start)
	for i := 0; i < workers; i++ {
		go func() {
			for {
				x := int(atomic.AddUint64(&current, 1))
				if x > end {
					wg.Done()
					return
				}

				fn(x - 1)
			}
		}()
	}

	wg.Wait()
}

// ParallelRanges divides up the the domain into segments and dispatches
// them to worker functions.
func ParallelRanges(start, end, workers int, fn func(start, end int)) {
	if end-start == 0 {
		return
	}
	if end-start < workers {
		fn(start, end)
		return
	}

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		a := i * start / end
		b := (i + 1) * start / end
		if i == workers-1 {
			b = end
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
