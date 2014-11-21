package assets

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const tempDirPrefix = "assets_TypeScript"

// TypeScript is a filter that runs the Microsoft TypeScript compiler on the
// assets. All files contained in the bundle must have a .ts extension. It is
// an error to use this filter on a bundle that contains a file without a .ts
// extension.
func TypeScript() Filter {
	return FilterFunc(func(bundle AssetBundle) (AssetBundle, error) {
		// Validate that the bundle contains only TypeScript files
		for _, fileName := range bundle.FileNames() {
			if filepath.Ext(fileName) != ".ts" {
				return nil, fmt.Errorf("`%s` does not end in .ts", fileName)
			}
		}

		// Create a temporary directory to compile into.
		dir, err := ioutil.TempDir(os.TempDir(), tempDirPrefix)
		if err != nil {
			return nil, err
		}
		// Remember to clean up after ourselves
		defer os.RemoveAll(dir)

		// Write out our files to the temporary directory.
		bundle, err = bundle.Filter(WriteToDir(dir))
		if err != nil {
			return nil, err
		}

		// Execute the TypeScript compiler.
		pathToCompiler, err := exec.LookPath("tsc")
		if err != nil {
			// There is no TypeScript compiler in the $PATH
			return nil, err
		}
		cmd := exec.Command(pathToCompiler, append([]string{"--outDir", dir}, bundle.FileNames()...)...)
		err = cmd.Run()
		if err != nil {
			return nil, err
		}

		// Create the new asset list for the resulting bundle.
		assets := make([]Asset, len(bundle.Assets()))
		for idx, fileName := range bundle.FileNames() {
			newFileName := strings.TrimSuffix(fileName, ".ts")
			newFileName = newFileName + ".js"

			contents, err := ioutil.ReadFile(filepath.Join(dir, newFileName))
			if err != nil {
				return nil, err
			}

			assets[idx] = NewAsset(
				newFileName,
				ioutil.NopCloser(bytes.NewReader(contents)),
			)
		}

		return &defaultBundle{
			currentName: bundle.Name(),
			assets:      assets,
		}, nil
	})
}
