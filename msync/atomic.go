package msync

import (
	"sync/atomic"
	"unsafe"
)

// Value is an implementation of atomic.Value which does not require consistent
// concrete types to be stored. It must not be copied after initialization.
type Value struct {
	v unsafe.Pointer
}

// Load returns the value set by the most recent Store.
func (v *Value) Load() (x interface{}) {
	ptr := (*interface{})(atomic.LoadPointer(&v.v))
	if ptr != nil {
		x = *ptr
	}

	return x
}

// Store sets the value of the Value to x.
func (v *Value) Store(x interface{}) {
	atomic.StorePointer(&v.v, unsafe.Pointer(&x))
}
