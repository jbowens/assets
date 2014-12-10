package assets

import (
	"os"
	"path/filepath"
	"regexp"
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
func (d *dir) Files(files ...string) (AssetBundle, error) {
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
	// an errorBundle that evaluates to an error.
	if err != nil {
		for _, a := range assets {
			if a.Contents() != nil {
				a.Contents().Close()
			}
		}
		return nil, err
	}

	return &defaultBundle{
		currentName: filepath.Base(d.Path),
		assets:      assets,
	}, nil
}

// MustFiles returns an AssetBundle containing only the files at the filenames
// given as arguments. If an error occurs, this function will panic.
func (d *dir) MustFiles(files ...string) AssetBundle {
	bundle, err := d.Files(files...)
	if err != nil {
		panic(err)
	}
	return bundle
}

// AllFiles returns an AssetBundle containing all the files in the directory.
func (d *dir) AllFiles() (AssetBundle, error) {
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
		return nil, err
	}

	return &defaultBundle{
		currentName: filepath.Base(d.Path),
		assets:      assets,
	}, nil
}

// MustAllFiles returns an AssetBundle containing all the files in the
// directory. If an error occurs, this function will panic.
func (d *dir) MustAllFiles() AssetBundle {
	bundle, err := d.AllFiles()
	if err != nil {
		panic(err)
	}
	return bundle
}

// Glob returns all files in the directory matching the glob expression.
func (d *dir) Glob(globExpr string) (AssetBundle, error) {
	globExpr = filepath.Join(d.Path, globExpr)

	filePaths, err := filepath.Glob(globExpr)
	if err != nil {
		return nil, err
	}

	assets := make([]Asset, len(filePaths))
	for idx, filePath := range filePaths {
		file, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}

		assets[idx] = &asset{
			fileName: filepath.Base(filePath),
			contents: file,
		}
	}

	return &defaultBundle{
		currentName: filepath.Base(d.Path),
		assets:      assets,
	}, nil
}

// Glob returns all files in the directory matching the glob expression.
// If an error occurs, this function will panic.
func (d *dir) MustGlob(globExpr string) AssetBundle {
	bundle, err := d.Glob(globExpr)
	if err != nil {
		panic(err)
	}
	return bundle
}

func (d *dir) WalkRegexp(regexp *regexp.Regexp) (AssetBundle, error) {
	assets := []Asset{}

	err := filepath.Walk(d.Path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && regexp.MatchString(info.Name()) {
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
		return nil, err
	}

	return &defaultBundle{
		currentName: filepath.Base(d.Path),
		assets:      assets,
	}, nil
}

func (d *dir) MustWalkRegexp(regexp *regexp.Regexp) AssetBundle {
	bundle, err := d.WalkRegexp(regexp)
	if err != nil {
		panic(err)
	}
	return bundle
}
