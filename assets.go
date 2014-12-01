package assets

import "io"

// AssetBundle holds a collection of assets. Filters may be applied to bundles,
// and the resulting assets outputted to disk.
type AssetBundle interface {
	// Name retrieves the name of the bundle.
	Name() string

	// Add adds a bundle of assets to this bundle.
	Add(...AssetBundle) AssetBundle

	// Filter performs the given filters on all assets contained within the
	// bundle. Filters are executed in the order they're received.
	Filter(...Filter) (AssetBundle, error)

	// MustFilter performs the given filters on all assets contained within the
	// bundle. Filters are executed in the order they're received. If an error
	// occurs in any filter, this function will panic.
	MustFilter(...Filter) AssetBundle

	// Assets returns all the assets contained within the bundle.
	Assets() []Asset

	// FileNames returns all of assets file names.
	FileNames() []string
}

// Directory represents a directory from which we can retrieve assets.
type Directory interface {
	// Files returns the provided files as a bundle.
	Files(...string) (AssetBundle, error)

	// MustFiles returns the provided files as a bundle. On error this
	// function will panic.
	MustFiles(...string) AssetBundle

	// AllFiles returns all files in the directory as a bundle.
	AllFiles() (AssetBundle, error)

	// MustAllFiles returns all files in the directory as a bundle. On error
	// this function will panic.
	MustAllFiles() AssetBundle

	// Glob returns all files in the directory matching the glob expression.
	Glob(string) (AssetBundle, error)

	// MustGlob returns all files in the directory matching the glob expression.
	// If an error occurs, this function will panic.
	MustGlob(string) (AssetBundle, error)
}

// Asset describes an individual asset file.
type Asset interface {
	FileName() string
	Contents() io.ReadCloser
	newCopy() (io.Reader, error)
}

// Filter defines a filter that can be applied to bundle of assets.
type Filter interface {
	// RunFilter takes a bundle, performs its operation on it, and
	// returns the resulting bundle
	RunFilter(AssetBundle) (AssetBundle, error)
}

// FilterFunc is an adapter type for wrapping simple functions in a Filter
// interface.
type FilterFunc func(AssetBundle) (AssetBundle, error)

// RunFilter executes the FilterFunc as a filter.
func (f FilterFunc) RunFilter(bundle AssetBundle) (AssetBundle, error) {
	return f(bundle)
}

// Bundle creates a new bundle with the given name.
func Bundle(name string) AssetBundle {
	return &defaultBundle{
		currentName: name,
		assets:      []Asset{},
	}
}
