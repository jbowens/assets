package assets

import (
	"io"
	"mime"
	"net/http"
	"path/filepath"
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
	mapOfAssets := make(map[string]Asset)
	mapOfTypes := make(map[string]string)

	for _, bundle := range bundles {
		for _, asset := range bundle.Assets() {
			mapOfAssets[asset.FileName()] = asset
			mimeType := mime.TypeByExtension(filepath.Ext(asset.FileName()))
			if mimeType != "" {
				mapOfTypes[asset.FileName()] = mimeType
			}
		}
	}

	return &servingMiddleware{
		pathPrefix: pathPrefix,
		assets:     mapOfAssets,
		types:      mapOfTypes,
	}
}

type servingMiddleware struct {
	pathPrefix string
	assets     map[string]Asset
	types      map[string]string
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

	mimeType, exists := m.types[fileName]
	if exists {
		w.Header().Set("Content-Type", mimeType)
	}

	io.Copy(w, asset.Contents())
}
