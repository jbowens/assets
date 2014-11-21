package assets

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func Fingerprint() Filter {

	hasher := md5.New()

	return FilterFunc(func(bundle AssetBundle) (AssetBundle, error) {
		assets := bundle.Assets()
		fingerprintedAssets := make([]Asset, len(assets))

		for idx, asset := range assets {
			buf := new(bytes.Buffer)
			_, err := buf.ReadFrom(asset.Contents())
			if err != nil {
				return nil, err
			}
			asset.Contents().Close()

			hasher.Reset()
			hasher.Write(buf.Bytes())
			hash := hex.EncodeToString(hasher.Sum(nil))

			ext := filepath.Ext(asset.FileName())
			base := strings.TrimSuffix(asset.FileName(), ext)
			filename := base + "-" + hash + ext

			fingerprintedAssets[idx] = NewAsset(filename, ioutil.NopCloser(buf))
		}

		return &defaultBundle{
			currentName: bundle.Name(),
			assets:      fingerprintedAssets,
		}, nil
	})
}
