package assets

// The Concat filter combines all provided assets into one single
// asset.
func Concat() Filter {
	return FilterFunc(func(bundle AssetBundle) (AssetBundle, error) {

		// TODO(jackson): Combine readers into MultiReader but with Close func

		return &defaultBundle{
			name:   bundle.Name(),
			assets: []Asset{},
		}, nil
	})
}
