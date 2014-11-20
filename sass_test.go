package assets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var simpleSassCompiled = `body {
  color: #333; }`

func TestSimpleSass(t *testing.T) {
	bundle := Dir("test_files/sass").Files("simple.scss").Filter(Sass())
	filenames, contents := bundleToFilenamesAndContents(t, bundle)

	assert.Contains(t, filenames, "simple.scss")
	assert.Contains(t, contents, simpleSassCompiled)
}
