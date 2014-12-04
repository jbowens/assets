package assets

// Prefix is a filter prepends every filename with the given prefix.
func Prefix(prefix string) Filter {
	return FilterFunc(func(bundle AssetBundle) (AssetBundle, error) {
		assets := bundle.Assets()
		prefixedAssets := make([]Asset, len(assets))

		for idx, asset := range assets {
			filename := prefix + asset.FileName()
			prefixedAssets[idx] = NewAsset(filename, asset.Contents())
		}

		return &defaultBundle{
			currentName: bundle.Name(),
			assets:      prefixedAssets,
		}, nil
	})
}
