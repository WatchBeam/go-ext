package msync

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValue(t *testing.T) {
	var v Value
	assert.Nil(t, v.Load())
	x := struct{ Foo int }{42}
	v.Store(x)
	assert.Equal(t, x, v.Load())
}

func TestValueRaces(t *testing.T) {
	// this'll cause the race detector to panic if it's unsafe
	var v Value
	for i := 0; i < 10; i++ {
		go func(i int) {
			v.Store(struct{ Foo int }{42})
			v.Load()
		}(i)
	}
}
