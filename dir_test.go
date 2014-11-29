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
	bundle, err := d.AllFiles()
	assert.Nil(t, err)
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
	bundle, err := d.Files("a", "c")
	assert.Nil(t, err)
	filenames, contents := bundleToFilenamesAndContents(t, bundle)

	assert.Contains(t, filenames, "a")
	assert.Contains(t, filenames, "c")
	assert.Contains(t, contents, "a")
	assert.Contains(t, contents, "c")
}

func TestDirGlob(t *testing.T) {
	d := Dir("test_files/javascript")
	bundle, err := d.Glob("*.js")
	assert.Nil(t, err)
	filenames, _ := bundleToFilenamesAndContents(t, bundle)

	assert.Len(t, filenames, 1)
	assert.Contains(t, filenames, "simple.js")
}
