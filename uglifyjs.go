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
	return &UglifyJSFilter{
		UseMinFileExtension: true,
		Mangle:              false,
		Compress:            false,
	}
}

// UglifyJSFilter is a filter that runs the uglifyjs command line utility
// over the assets. It has a few options to specify flags to the cli.
type UglifyJSFilter struct {
	UseMinFileExtension bool
	Mangle              bool
	Compress            bool
}

// RunFilter executes the UglifyJS filter on the given bundle.
func (u *UglifyJSFilter) RunFilter(bundle AssetBundle) (AssetBundle, error) {
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

func (u *UglifyJSFilter) uglify(source io.ReadCloser) (io.ReadCloser, error) {
	cmd := exec.Command("uglifyjs", u.commandArgs()...)
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

func (u *UglifyJSFilter) commandArgs() []string {
	args := []string{}

	if u.Mangle {
		args = append(args, "--mangle")
	}

	if u.Compress {
		args = append(args, "--compress")
	}

	args = append(args, "-")
	return args
}

func (u *UglifyJSFilter) fileName(old string) string {
	if !u.UseMinFileExtension || filepath.Ext(old) != ".js" {
		return old
	}

	newName := strings.TrimSuffix(old, ".js")
	newName = newName + ".min.js"
	return newName
}
