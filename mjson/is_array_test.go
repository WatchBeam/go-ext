package mjson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsArray(t *testing.T) {
	assert.True(t, IsArray([]byte(`[42]`)))
	assert.True(t, IsArray([]byte(`  [42]  `)))
	assert.False(t, IsArray([]byte(`{}`)))
	assert.False(t, IsArray([]byte(`null`)))
	assert.False(t, IsArray([]byte(`{`)))
	assert.False(t, IsArray([]byte(`42`)))
}
