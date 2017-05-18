package msync

import (
	"runtime"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func parallelTester(max, workers int, ch chan<- int) {
	Parallel(0, max, workers, func(x int) {
		ch <- x
	})
}
func parallelRangesTester(max, workers int, ch chan<- int) {
	ParallelRanges(0, max, workers, func(start, end int) {
		for i := start; i < end; i++ {
			ch <- i
		}
	})
}

func testRangeGenerator(t *testing.T, fn func(max, workers int, ch chan<- int)) {
	for max := 0; max < 100; max++ {
		var expected []int
		for i := 0; i < max; i++ {
			expected = append(expected, i)
		}

		for workers := 1; workers < 5; workers++ {
			ch := make(chan int, max)
			fn(max, workers, ch)
			close(ch)

			var actual []int
			for n := range ch {
				actual = append(actual, n)
			}

			sort.Sort(sort.IntSlice(actual))
			assert.Equal(t, expected, actual)
		}
	}
}

func benchmarkRangeGenerator(b *testing.B, fn func(max, workers int, ch chan<- int)) {
	ch := make(chan int, b.N)
	b.ResetTimer()
	fn(b.N, runtime.GOMAXPROCS(0), ch)
}

func TestParallel(t *testing.T)      { testRangeGenerator(t, parallelTester) }
func TestParallelRange(t *testing.T) { testRangeGenerator(t, parallelRangesTester) }

func BenchmarkParallel(b *testing.B)      { benchmarkRangeGenerator(b, parallelTester) }
func BenchmarkParallelRange(b *testing.B) { benchmarkRangeGenerator(b, parallelRangesTester) }
