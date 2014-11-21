package assets

import (
	"os"
	"path/filepath"
)

type dir struct {
	Path string
}

// Dir creates a Directory representing the directory at the given path.
func Dir(path string) Directory {
	return &dir{
		Path: path,
	}
}

// Files returns an AssetBundle containing only the files at the filenames
// given as arguments.
func (d *dir) Files(files ...string) AssetBundle {
	assets := make([]Asset, len(files))

	var err error
	for idx, fileName := range files {
		var file *os.File
		file, err = os.Open(filepath.Join(d.Path, fileName))

		assets[idx] = &asset{
			fileName: fileName,
			contents: file,
		}
	}

	// If an error occurred, close any readers we opened and return
	// an ErrorBundle that evaluates to an error.
	if err != nil {
		for _, a := range assets {
			if a.Contents() != nil {
				a.Contents().Close()
			}
		}
		return &ErrorBundle{err: err}
	}

	return &defaultBundle{
		currentName: filepath.Base(d.Path),
		assets:      assets,
	}
}

// AllFiles returns an AssetBundle containing all the files in the directory.
func (d *dir) AllFiles() AssetBundle {
	assets := []Asset{}

	err := filepath.Walk(d.Path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				file, err := os.Open(path)
				if err != nil {
					return err
				}

				assets = append(assets, &asset{
					fileName: info.Name(),
					contents: file,
				})
			}
			return nil
		})

	if err != nil {
		return &ErrorBundle{err: err}
	}

	return &defaultBundle{
		currentName: filepath.Base(d.Path),
		assets:      assets,
	}
}
