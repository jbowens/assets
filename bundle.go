package assets

// defaultBundle is the default implementation of the AssetBundle interface.
type defaultBundle struct {
	// currentName holds the name of the bundle
	currentName string

	// assets contains a slice of all the assets in the system.
	assets []Asset
}

func (b *defaultBundle) Add(otherBundles ...AssetBundle) AssetBundle {
	// TODO(jackson): Fix the order of added assets
	for _, bundle := range otherBundles {
		b.assets = append(b.assets, bundle.Assets()...)
	}
	return b
}

func (b *defaultBundle) Assets() []Asset {
	return b.assets
}

func (b *defaultBundle) Filter(filters ...Filter) (AssetBundle, error) {
	var bundle AssetBundle = b
	var err error
	for _, f := range filters {
		bundle, err = f.RunFilter(bundle)
		if err != nil {
			return nil, err
		}
	}

	return bundle, nil
}

func (b *defaultBundle) MustFilter(filters ...Filter) AssetBundle {
	bundle, err := b.Filter(filters...)
	if err != nil {
		panic(err)
	}
	return bundle
}

func (b *defaultBundle) Name() string {
	return b.currentName
}

func (b *defaultBundle) FileNames() []string {
	names := make([]string, len(b.assets))

	for idx, asset := range b.assets {
		names[idx] = asset.FileName()
	}

	return names
}
