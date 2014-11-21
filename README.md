assets
======
[![GoDoc](https://godoc.org/github.com/jbowens/assets?status.svg)](https://godoc.org/github.com/jbowens/assets)

Assets management for golang web applications

## Examples

```go
coreJS, err := assets.Dir("assets/js")
  .Files("core.js",
         "util.js",
         "models.js",
         "network.js")

if err != nil {
    return err
}

coreJS, err = coreJS.Filter(
  assets.Concat(),
  assets.Uglify(),
  assets.WriteToDir("generated/assets/js")
  ).DependsOn(jqueryBundle)
```

In some places, assets offers Must functions that will panic on error.
```go
homepageCSS := assets.Dir("assets/css/home").MustAllFiles()
  .MustFilter(assets.Concat(), assets.Sass(), assets.WriteToDir("generated/assets/css"))
```

Multiple filters may be combined into one.
```go
cssPipeline := assets.Combine(assets.Concat(), assets.Sass(), assets.WriteToDir("gen"))

homepage := assets.Dir("assets/homepage").MustAllFiles().MustFilter(cssPipeline)
widgets := assets.Dir("assets/widgets").MustAllFiles().MustFilter(cssPipeline)
```
