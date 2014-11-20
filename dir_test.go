package assets

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func bundleToFilenamesAndContents(t *testing.T, bundle AssetBundle) ([]string, []string) {
	filenames := []string{}
	contents := []string{}
	for _, a := range bundle.Assets() {
		filenames = append(filenames, a.FileName())
		buf := new(bytes.Buffer)
		_, err := buf.ReadFrom(a.Contents())
		if err != nil {
			t.Fatal(err)
		}
		contents = append(contents, strings.TrimSpace(buf.String()))
	}
	return filenames, contents
}

func TestDirAllFiles(t *testing.T) {
	d := Dir("test_files/abc")
	bundle := d.AllFiles()
	filenames, contents := bundleToFilenamesAndContents(t, bundle)

	assert.Contains(t, filenames, "a")
	assert.Contains(t, filenames, "b")
	assert.Contains(t, filenames, "c")
	assert.Contains(t, contents, "a")
	assert.Contains(t, contents, "b")
	assert.Contains(t, contents, "c")
}

func TestDirFiles(t *testing.T) {
	d := Dir("test_files/abc")
	bundle := d.Files("a", "c")
	filenames, contents := bundleToFilenamesAndContents(t, bundle)

	assert.Contains(t, filenames, "a")
	assert.Contains(t, filenames, "c")
	assert.Contains(t, contents, "a")
	assert.Contains(t, contents, "c")
}
