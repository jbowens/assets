package assets

import (
	"bytes"
	"io"
	"io/ioutil"
)

// asset is the default Asset implementation
type asset struct {
	fileName string
	contents io.ReadCloser
}

// NewAsset creates an Asset with the given filename and ReadCloser.
func NewAsset(fileName string, contents io.ReadCloser) Asset {
	return &asset{
		fileName: fileName,
		contents: contents,
	}
}

// FileName retrieves the current file name of the asset.
func (a *asset) FileName() string {
	return a.fileName
}

// Contents returns a ReadCloser for the contents of the asset.
func (a *asset) Contents() io.ReadCloser {
	return a.contents
}

func (a *asset) newCopy() (io.Reader, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(a.contents)
	if err != nil {
		return nil, err
	}
	a.Contents().Close()
	a.contents = ioutil.NopCloser(buf)
	return bytes.NewReader(buf.Bytes()), nil
}
