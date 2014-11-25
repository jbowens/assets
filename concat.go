package assets

import (
	"io"
	"path/filepath"
)

// Concat combines all provided assets into one asset. The resulting
// AssetBundle will contain a single asset.
func Concat() Filter {
	return FilterFunc(func(bundle AssetBundle) (AssetBundle, error) {
		assets := bundle.Assets()
		readClosers := make([]io.ReadCloser, len(assets))

		var extensions = make(map[string]int)
		for idx, asset := range assets {
			readClosers[idx] = asset.Contents()
			extensions[filepath.Ext(asset.FileName())]++
		}

		var mostCommonExt string
		for ext, count := range extensions {
			if count > extensions[mostCommonExt] {
				mostCommonExt = ext
			}
		}

		return &defaultBundle{
			currentName: bundle.Name(),
			assets: []Asset{
				&asset{
					fileName: bundle.Name() + mostCommonExt,
					contents: newMultiReadCloser(readClosers...),
				},
			},
		}, nil
	})
}
