package assets

// errorBundle is an implementation of an AssetBundle for a pipeline
// that has resulted in an error.
type errorBundle struct {
	err error
}

// Add adds a bundle of assets to this bundle.
func (b *errorBundle) Add(otherBundle AssetBundle) AssetBundle {
	return b
}

// Assets returns all the assets contained within the bundle.
func (b *errorBundle) Assets() []Asset {
	return []Asset{}
}

// Filter performs the given filters on all assets contained within the
// bundle. Filters are executed in the order they're received.
func (b *errorBundle) Filter(filters ...Filter) AssetBundle {
	return b
}

// Name retrieves the name of the bundle.
func (b *errorBundle) Name() string {
	return "error"
}

// MustWrite writes out all assets in the bundle. It panics if an error
// occurred during the pipeline.
func (b *errorBundle) MustWrite(dir string) AssetBundle {
	panic(b.err)
}

// Write writes out all assets in the bundle to the provided directory.
func (b *errorBundle) Write(dir string) (AssetBundle, error) {
	return nil, b.err
}
