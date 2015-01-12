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
		// If this an empty bundle, just skip this.
		if len(bundle.FileNames()) == 0 {
			return bundle, nil
		}

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

		args := append([]string{"--outDir", dir}, bundle.FileNames()...)
		for i := 2; i < len(args); i++ {
			args[i] = filepath.Join(dir, args[i])
		}

		cmd := exec.Command(pathToCompiler, args...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return nil, fmt.Errorf("Error running tsc: %v, output: %v", err, string(output))
		}

		// Create the new asset list for the resulting bundle.
		var assets []Asset
		for _, fileName := range bundle.FileNames() {
			if strings.HasSuffix(fileName, ".d.ts") {
				continue
			}

			newFileName := strings.TrimSuffix(fileName, ".ts")
			newFileName = newFileName + ".js"

			contents, err := ioutil.ReadFile(filepath.Join(dir, newFileName))
			if err != nil {
				return nil, err
			}

			assets = append(assets, NewAsset(
				newFileName,
				ioutil.NopCloser(bytes.NewReader(contents)),
			))
		}

		return &defaultBundle{
			currentName: bundle.Name(),
			assets:      assets,
		}, nil
	})
}
