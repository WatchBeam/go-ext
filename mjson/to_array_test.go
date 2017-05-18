package mjson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToArray(t *testing.T) {
	cases := []struct {
		items    []string
		expected string
	}{
		{[]string{}, "[]"},
		{[]string{"1"}, `[1]`},
		{[]string{"1", "2"}, `[1,2]`},
		{[]string{"1", "2", "3"}, `[1,2,3]`},
	}

	for _, tc := range cases {
		asBytes := [][]byte{}
		for _, item := range tc.items {
			asBytes = append(asBytes, []byte(item))
		}
		assert.Equal(t, tc.expected, string(ToArray(asBytes)))
	}
}
