package merr

import (
	"errors"
	"testing"

	perror "github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestComposedEmpty(t *testing.T) {
	e := Compose()
	assert.True(t, e.Empty())
	e.Add(errors.New("oh no!"))
	assert.False(t, e.Empty())
}

func TestComposedPrintsSimpleMessage(t *testing.T) {
	e := Compose(errors.New("oh no!"))
	assert.Equal(t, "error #0: oh no!", e.Error())
}

func TestComposedPrintsStackTrace(t *testing.T) {
	e := Compose(perror.New("oh no!"))
	assert.Regexp(t, "error #0: oh no!\n.+merr\\.TestComposedPrintsStackTrace", e.Error())
}

func TestConcurrentComposedEmpty(t *testing.T) {
	e := ComposeConcurrent()
	assert.True(t, e.Empty())
	e.Add(errors.New("oh no!"))
	assert.False(t, e.Empty())
}

func TestConcurrentComposePrintsSimpleMessage(t *testing.T) {
	e := ComposeConcurrent(errors.New("oh no!"))
	assert.Equal(t, "error #0: oh no!", e.Error())
}

func TestConcurrentComposePrintsStackTrace(t *testing.T) {
	e := ComposeConcurrent(perror.New("oh no!"))
	assert.Regexp(t, "error #0: oh no!\n.+merr\\.TestConcurrentComposePrintsStackTrace", e.Error())
}
