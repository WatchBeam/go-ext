package msync

import (
	"sync/atomic"
	"time"
)

// ThrottleTail returns a function that calls fn() at most every duration, when
// the first call is on the tail of the duration.
func ThrottleTail(duration time.Duration, fn func()) func() {
	var lastCall int64
	return func() {
		last := atomic.LoadInt64(&lastCall)
		now := time.Now().UnixNano()
		if last > now {
			return
		}

		if atomic.CompareAndSwapInt64(&lastCall, last, now+int64(duration)) {
			go func() {
				time.Sleep(duration)
				fn()
			}()
		}
	}
}

// ThrottleTail returns a function that calls fn() at most every duration, when
// the first call is on the head of the duration.
func ThrottleHead(duration time.Duration, fn func()) func() {
	var lastCall int64
	return func() {
		last := atomic.LoadInt64(&lastCall)
		now := time.Now().UnixNano()
		if last+int64(duration) > now {
			return
		}

		if atomic.CompareAndSwapInt64(&lastCall, last, now) {
			fn()
		}
	}
}
