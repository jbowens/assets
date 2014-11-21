package assets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var simpleSassCompiled = `body {
  color: #333; }`

func TestSimpleSass(t *testing.T) {
	bundle, err := Dir("test_files/sass").MustFiles("simple.scss").Filter(Sass())
	assert.Nil(t, err)
	filenames, contents := bundleToFilenamesAndContents(t, bundle)

	assert.Contains(t, filenames, "simple.css")
	assert.Contains(t, contents, simpleSassCompiled)
}
