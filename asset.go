package assets

import "io"

// asset is the default Asset implementation
type asset struct {
	fileName string
	contents io.ReadCloser
}

func NewAsset(fileName string, contents io.ReadCloser) Asset {
	return &asset{
		fileName: fileName,
		contents: contents,
	}
}

func (a *asset) FileName() string {
	return a.fileName
}

func (a *asset) Contents() io.ReadCloser {
	return a.contents
}
