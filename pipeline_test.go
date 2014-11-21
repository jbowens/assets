package assets

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func expectFile(t *testing.T, path, contents string) {
	bytes, err := ioutil.ReadFile(path)
	assert.Nil(t, err)

	fileContents := string(bytes)
	assert.Equal(t, contents, fileContents)
}

func TestSimplePipeline(t *testing.T) {
	bundle := Dir("test_files/abc").AllFiles().
		Filter(Concat(), WriteToDir("test_files/generated"))

	//assert.Nil(t, err)
	assert.NotNil(t, bundle)
	expectFile(t, "test_files/generated/abc", "a\nb\nc\n")
}

func TestSimpleFingerprintingPipeline(t *testing.T) {
	bundle := Dir("test_files/css/").AllFiles().
		Filter(Fingerprint(), WriteToDir("test_files/generated"))

	// TODO(jackson): Clean up these tests once I figure out how to propagate
	// errors cleanly.

	//assert.Nil(t, err)
	assert.NotNil(t, bundle)
}
