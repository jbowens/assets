package assets

// defaultBundle is the default implementation of the AssetBundle interface.
type defaultBundle struct {
	// Name stores the name of the bundle, usually used for outputted files.
	name string

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

func (b *defaultBundle) Filter(filters ...Filter) AssetBundle {
	// TODO: Implement
	return nil
}

func (b *defaultBundle) Name() string {
	return b.name
}
