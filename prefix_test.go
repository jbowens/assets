package assets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixFilter(t *testing.T) {
	d := Dir("test_files/abc")
	bundle, err := d.AllFiles()
	assert.Nil(t, err)
	bundle, err = bundle.Filter(Prefix("prefix_"))
	assert.Nil(t, err)
	filenames, contents := bundleToFilenamesAndContents(t, bundle)

	assert.Contains(t, filenames, "prefix_a")
	assert.Contains(t, filenames, "prefix_b")
	assert.Contains(t, filenames, "prefix_c")
	assert.Contains(t, contents, "a")
	assert.Contains(t, contents, "b")
	assert.Contains(t, contents, "c")
}
