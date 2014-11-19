package assets

type dir struct{}

func Dir(path string) *bundle {

	return &bundle{Name: path}
}
