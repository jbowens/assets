package assets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimplePipeline(t *testing.T) {
	bundle, err := Dir("test_files/abc").AllFiles().
		Filter(Concat()).Write("test_files/generated")

	assert.Nil(t, err)
	assert.NotNil(t, bundle)
}
