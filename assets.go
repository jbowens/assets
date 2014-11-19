package assets

// Expected use case:
//
// assets.Bundle("core").Dir("assets")
//   .Files("core.js",
//          "util.js",
//          "models.js",
//          "network.js")
//   .Filter(assets.Combine(), assets.Minify())
//   .DependsOn(dependencyBundle)
//

// Bundle holds a collection of assets. Filters may be applied to bundles,
// and the resulting assets outputted to disk.
type AssetBundle interface {

	// Dir returns an AssetBundle chdir'd into the specified directory.
	Dir(string) AssetBundle

	// Value returns the resulting, evaluated AssetBundle or an error,
	// depending on the result of evaluation.
	Value() (AssetBundle, error)

	// Must returns the resulting AssetBundle. If an error resulted during
	// evaluation, Must will panic.
	Must() AssetBundle
}

// Directory represents a directory from which we can retrieve assets.
type Directory interface {
	// Files returns the provided files as a bundle.
	Files(...string) AssetBundle

	// AllFiles returns all files in the directory as a bundle
	AllFiles() AssetBundle
}

// Assets describes an individual asset file.
type Asset interface {
	FileName() string
	Contents() []byte
}

// Filter defines a filter that can be applied to bundle of assets.
type Filter interface {
	// RunFilter takes a slice of assets, performs its operation on them, and
	// returns the resulting asset slice.
	RunFilter([]Asset) ([]Asset, error)
}

// Bundle creates a new bundle with the given name.
func Bundle(name string) AssetBundle {
	return &bundle{
		Name: name,
	}
}
