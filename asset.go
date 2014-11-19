package assets

// asset is the default Asset implementation
type asset struct {
	fileName string
	contents []byte
}

func NewAsset(fileName string, contents []byte) Asset {
	return &asset{
		fileName: fileName,
		contents: contents,
	}
}

func (a *asset) FileName() string {
	return a.fileName
}

func (a *asset) Contents() []byte {
	return a.contents
}
