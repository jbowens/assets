assets
======
[![GoDoc](https://godoc.org/github.com/jbowens/assets?status.svg)](https://godoc.org/github.com/jbowens/assets) [![Build Status](https://travis-ci.org/jbowens/assets.svg?branch=master)](https://travis-ci.org/jbowens/assets)

Assets provides general assets management for golang web applications.

## Examples

### Managing dependencies

Dependencies can be added to a bundle at any point in the pipeline.

Here's an example of compiling core javascript together in a single file,
homepage-specific javascript into a single file, and then including both
into a homepage.

```go
var coreJS assets.AssetBundle
var homeJS assets.AssetBundle

func onStartUp() {

	// On start up, compile your assets.
	pipeline := []assets.Filter{
		assets.Concat(),
		assets.UglifyJS(),
		assets.Fingerprint(),
		assets.WriteToDir("generated"),
	}

	coreJS = assets.Dir("assets/js/core").MustAllFiles().MustFilter(pipeline...)
	homeJS = assets.Dir("assets/js/home").MustAllFiles().MustFilter(pipeline...).Add(coreJS)
}

func whenRenderingHome() {
	// When including your javascript, include all files listed by
	// homeJS.FileNames().
	for _, file := range homeJS.FileNames() {
		fmt.Println(file)
	}
}
```

This example will print:
```
home-a00e289d7a1520aa1a2e824b404dc8be.min.js
core-9627c435b535b37cb247355c8b36f9aa.min.js
```

## Credits

This package was inspired by [ghemawat/stream](https://github.com/ghemawat/stream).

It was written by Jackson Owens.
