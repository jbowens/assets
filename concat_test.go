package assets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcatEndToEnd(t *testing.T) {
	d := Dir("test_files/abc")
	bundle, err := d.AllFiles()
	assert.Nil(t, err)
	bundle = bundle.Filter(Concat())
	filenames, contents := bundleToFilenamesAndContents(t, bundle)

	assert.Contains(t, filenames, "abc")
	assert.Contains(t, contents, "a\nb\nc")
}
