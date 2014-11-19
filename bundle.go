package assets

// bundle is the default implementation of the AssetBundle interface.
type bundle struct {
	// Name stores the name of the bundle, usually used for outputted files.
	Name string

	// assets contains a slice of all the assets in the system.
	assets []Asset
}

func (b *bundle) Add(otherBundle AssetBundle) AssetBundle {
	b.assets = append(b.assets, otherBundle.Assets()...)
	return b
}

func (b *bundle) Assets() []Asset {
	return b.assets
}

func (b *bundle) Filter(filters ...Filter) AssetBundle {
	// TODO: Implement
	return nil
}

func (b *bundle) Value() (AssetBundle, error) {
	// TODO: Implement
	return nil, nil
}

func (b *bundle) Must() AssetBundle {
	val, err := b.Value()
	if err != nil {
		panic(err)
	}
	return val
}
