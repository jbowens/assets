package assets

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/codegangsta/negroni"
)

// MemoryServingMiddleware returns a negroni middleware that will serve all
// assets in the given bundles for all matching requests whose path begins
// with pathPrefix.
//
// If you have large assets, do not use this middleware. In production you
// should use nginx to deliver assets.
func MemoryServingMiddleware(pathPrefix string, bundles ...AssetBundle) negroni.Handler {
	mapOfAssets := make(map[string][]byte)

	for _, bundle := range bundles {
		for _, asset := range bundle.Assets() {
			bytes, err := ioutil.ReadAll(asset.Contents())
			if err != nil {
				mapOfAssets[asset.FileName()] = bytes
			}
		}
	}

	return &servingMiddleware{
		pathPrefix: pathPrefix,
		assets:     mapOfAssets,
	}
}

type servingMiddleware struct {
	pathPrefix string
	assets     map[string][]byte
}

func (m *servingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	if !strings.HasPrefix(r.URL.Path, m.pathPrefix) {
		next(w, r)
		return
	}

	fileName := strings.TrimPrefix(r.URL.Path, m.pathPrefix)
	asset, exists := m.assets[fileName]

	if !exists {
		next(w, r)
		return
	}

	w.Write(asset)
}
