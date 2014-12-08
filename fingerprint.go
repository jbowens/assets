package assets

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"strings"
)

// Fingerprint is a filter that sets asset filenames to include a md5 hash
// of the file contents. This can help circumvent browser caching when changes
// are made.
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

			fileNamePieces := strings.SplitN(asset.FileName(), ".", 2)
			var filename string
			if len(fileNamePieces) == 2 {
				filename = fileNamePieces[0] + "-" + hash + "." + fileNamePieces[1]
			} else {
				filename = fileNamePieces[0] + "-" + hash
			}

			fingerprintedAssets[idx] = NewAsset(filename, ioutil.NopCloser(buf))
		}

		return &defaultBundle{
			currentName: bundle.Name(),
			assets:      fingerprintedAssets,
		}, nil
	})
}
