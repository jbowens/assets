package assets

import "io"

// Concat combines all provided assets into one asset. The resulting
// AssetBundle will contain a single asset.
func Concat() Filter {
	return FilterFunc(func(bundle AssetBundle) (AssetBundle, error) {
		assets := bundle.Assets()
		readClosers := make([]io.ReadCloser, len(assets))
		for idx, asset := range assets {
			readClosers[idx] = asset.Contents()
		}

		return &defaultBundle{
			currentName: bundle.Name(),
			assets: []Asset{
				&asset{
					fileName: bundle.Name(),
					contents: MultiReadCloser(readClosers...),
				},
			},
		}, nil
	})
}
