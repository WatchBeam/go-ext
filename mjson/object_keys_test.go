package mjson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var keysFixture = []byte(`{"foo":42,"bar":{},"wut":null}`)

func TestObjectKeys(t *testing.T) {
	keys, err := Keys(keysFixture)
	assert.Nil(t, err)
	assert.Equal(t, []string{"foo", "bar", "wut"}, keys)

	keys, err = Keys([]byte(`null`))
	assert.Nil(t, err)
	assert.Equal(t, []string{}, keys)
}

func BenchmarkKeys(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Keys(keysFixture)
	}
}
