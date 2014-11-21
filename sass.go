package assets

import (
	"bytes"
	"io/ioutil"
	"strings"

	sass "github.com/suapapa/go_sass"
)

func Sass() Filter {
	var compiler sass.Compiler

	return FilterFunc(func(bundle AssetBundle) (AssetBundle, error) {
		assets := bundle.Assets()
		compiledAssets := make([]Asset, len(assets))
		for idx, asset := range assets {
			var err error

			buf := new(bytes.Buffer)
			_, err = buf.ReadFrom(asset.Contents())
			if err != nil {
				return nil, err
			}
			asset.Contents().Close()

			compiledString, err := compiler.Compile(buf.String())
			if err != nil {
				return nil, err
			}
			compiledReader := ioutil.NopCloser(strings.NewReader(compiledString))

			filename := asset.FileName()
			if strings.HasSuffix(filename, ".scss") ||
				strings.HasSuffix(filename, ".sass") {
				filename = strings.TrimSuffix(filename, ".scss")
				filename = strings.TrimSuffix(filename, ".sass")
				filename = filename + ".css"
			}

			compiledAssets[idx] = NewAsset(filename, compiledReader)
		}

		return &defaultBundle{
			currentName: bundle.Name(),
			assets:      compiledAssets,
		}, nil
	})
}
