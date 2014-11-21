assets
======
[![GoDoc](https://godoc.org/github.com/jbowens/assets?status.svg)](https://godoc.org/github.com/jbowens/assets)

Assets management for golang web applications

## Example use cases

```go
core_js := assets.Dir("assets/js")
  .Files("core.js",
         "util.js",
         "models.js",
         "network.js")
  .Filter(assets.Concat(), assets.Uglify(), WriteToDir("generated/assets/js"))
  .DependsOn(jqueryBundle)
```

```go
homepage_css := assets.Dir("assets/css/home").AllFiles()
  .Filter(assets.Concat(), assets.Sass(), WriteToDir("generated/assets/css")
```
