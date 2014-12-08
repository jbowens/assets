package assets

import (
	"bytes"
	"io"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strings"
)

// UglifyJS is a filter that runs the uglifyjs command line utility over
// assets. By default, it replaces '.js' extensions with '.min.js'.
func UglifyJS() Filter {
	return &uglifyJS{
		useMinFileExtension: true,
	}
}

type uglifyJS struct {
	useMinFileExtension bool
}

func (u *uglifyJS) RunFilter(bundle AssetBundle) (AssetBundle, error) {
	// Make sure uglifyjs is in the PATH
	_, err := exec.LookPath("uglifyjs")
	if err != nil {
		return nil, err
	}

	assets := bundle.Assets()
	uglifiedAssets := make([]Asset, len(assets))

	for idx, asset := range assets {
		uglified, err := u.uglify(asset.Contents())
		if err != nil {
			return nil, err
		}
		uglifiedAssets[idx] = NewAsset(u.fileName(asset.FileName()), uglified)
	}

	return &defaultBundle{
		currentName: bundle.Name(),
		assets:      uglifiedAssets,
	}, nil
}

func (u *uglifyJS) uglify(source io.ReadCloser) (io.ReadCloser, error) {
	cmd := exec.Command("uglifyjs", "--mangle", "-")
	stdIn, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(stdIn, source)
	if err != nil {
		return nil, err
	}
	source.Close()
	stdIn.Close()

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return ioutil.NopCloser(bytes.NewReader(output)), nil
}

func (u *uglifyJS) fileName(old string) string {
	if !u.useMinFileExtension || filepath.Ext(old) != ".js" {
		return old
	}

	newName := strings.TrimSuffix(old, ".js")
	newName = newName + ".min.js"
	return newName
}
