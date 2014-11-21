package assets

import (
	"io"
	"os"
	"path/filepath"
)

// WriteToDir writes the bundle's assets out to the directory passed as an
// argument. It leaves the bundle's assets unmodified.
func WriteToDir(directory string) Filter {
	return FilterFunc(func(bundle AssetBundle) (AssetBundle, error) {
		assets := bundle.Assets()
		for _, asset := range assets {
			fileName := filepath.Join(directory, asset.FileName())

			err := os.MkdirAll(filepath.Dir(fileName), os.ModePerm)
			if err != nil {
				return nil, err
			}

			f, err := os.Create(fileName)
			if err != nil {
				return nil, err
			}

			contents, err := asset.newCopy()
			if err != nil {
				return nil, err
			}

			_, err = io.Copy(f, contents)
			if err != nil {
				return nil, err
			}
		}

		return bundle, nil
	})
}
