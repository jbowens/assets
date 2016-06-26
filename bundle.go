package assets

import "net/http"

// Bundle represents a set of assets. It provides access to the current
// relative paths that the assets are served from and can be used as a
// http.Handler that will serve the assets at those same paths.
type Bundle interface {
	http.Handler
	RelativePaths() []string
}
