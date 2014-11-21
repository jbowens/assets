assets
======
[![GoDoc](https://godoc.org/github.com/jbowens/assets?status.svg)](https://godoc.org/github.com/jbowens/assets)

Assets management for golang web applications

## Examples

```go
core_js, err := assets.Dir("assets/js")
  .Files("core.js",
         "util.js",
         "models.js",
         "network.js")

if err != nil {
    return err
}

core_js, err = core_js.Filter(
  assets.Concat(),
  assets.Uglify(),
  WriteToDir("generated/assets/js")
  ).DependsOn(jqueryBundle)
```

In some places, assets offers Must functions that will panic on error.
```go
homepage_css := assets.Dir("assets/css/home").MustAllFiles()
  .MustFilter(assets.Concat(), assets.Sass(), WriteToDir("generated/assets/css")
```
