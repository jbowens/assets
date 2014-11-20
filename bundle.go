package assets

import (
	"io"
	"os"
	"path/filepath"
)

// defaultBundle is the default implementation of the AssetBundle interface.
type defaultBundle struct {
	// currentName holds the name of the bundle
	currentName string

	// assets contains a slice of all the assets in the system.
	assets []Asset
}

func (b *defaultBundle) Add(otherBundle AssetBundle) AssetBundle {
	b.assets = append(b.assets, otherBundle.Assets()...)
	return b
}

func (b *defaultBundle) Assets() []Asset {
	return b.assets
}

func (b *defaultBundle) Filter(filters ...Filter) AssetBundle {

	var bundle AssetBundle = b
	var err error
	for _, f := range filters {
		bundle, err = f.RunFilter(bundle)
		if err != nil {
			return &ErrorBundle{err: err}
		}
	}

	return bundle
}

func (b *defaultBundle) Name() string {
	return b.currentName
}

func (b *defaultBundle) MustWrite(dir string) AssetBundle {
	bundle, err := b.Write(dir)
	if err != nil {
		panic(err)
	}
	return bundle
}

func (b *defaultBundle) Write(dir string) (AssetBundle, error) {
	for _, asset := range b.assets {
		fileName := filepath.Join(dir, asset.FileName())
		f, err := os.Create(fileName)
		if err != nil {
			return nil, err
		}

		_, err = io.Copy(f, asset.Contents())
		if err != nil {
			return nil, err
		}

		asset.Contents().Close()
	}

	return b, nil
}
