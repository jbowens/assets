package assets

// Combine returns a Filter that is the concatenation of all filter arguments.
func Combine(filters ...Filter) Filter {
	return FilterFunc(func(bundle AssetBundle) (AssetBundle, error) {
		var err error

		for _, f := range filters {
			bundle, err = f.RunFilter(bundle)
			if err != nil {
				return nil, err
			}
		}

		return bundle, err
	})
}
