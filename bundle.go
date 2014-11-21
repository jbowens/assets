package assets

// defaultBundle is the default implementation of the AssetBundle interface.
type defaultBundle struct {
	// currentName holds the name of the bundle
	currentName string

	// assets contains a slice of all the assets in the system.
	assets []Asset
}

func (b *defaultBundle) Add(otherBundle AssetBundle) AssetBundle {
	b.assets = append(b.assets, otherBundle.Assets()...)
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
