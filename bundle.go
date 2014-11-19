package assets

// bundle is the default implementation of the AssetBundle interface.
type bundle struct {
	// Name stores the name of the bundle, usually used for outputted files.
	Name string
}

func (b *bundle) Value() (*bundle, error) {
	// TODO: Implement
	return nil, nil
}

func (b *bundle) Must() *bundle {
	val, err := b.Value()
	if err != nil {
		panic(err)
	}
	return val
}
