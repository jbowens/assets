package assets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleSass(t *testing.T) {
	bundle, err := Dir("test_files/sass").MustFiles("simple.scss").Filter(Sass())
	assert.Nil(t, err)
	filenames, _ := bundleToFilenamesAndContents(t, bundle)

	assert.Contains(t, filenames, "simple.css")
}
