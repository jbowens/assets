package assets

// ErrorBundle is an implementation of an AssetBundle for a pipeline
// that has resulted in an error.
type ErrorBundle struct {
	err error
}

func (b *ErrorBundle) Add(otherBundle AssetBundle) AssetBundle {
	return b
}

func (b *ErrorBundle) Assets() []Asset {
	return []Asset{}
}

func (b *ErrorBundle) Filter(filters ...Filter) AssetBundle {
	return b
}
