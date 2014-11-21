package assets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleUglifyJS(t *testing.T) {
	bundle, err := Dir("test_files/javascript").MustFiles("simple.js").Filter(UglifyJS())
	assert.Nil(t, err)
	filenames, contents := bundleToFilenamesAndContents(t, bundle)

	assert.Contains(t, filenames, "simple.min.js")
	assert.Contains(t, contents, "var x=5;")
}
