package assets

type dir struct{}

func (b *bundle) Dir(path string) *dir {

	return &dir{}
}
