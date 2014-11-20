assets
======

Assets management for golang web applications

## Example use cases

```go
core_js := assets.Dir("assets/js")
  .Files("core.js",
         "util.js",
         "models.js",
         "network.js")
  .Filter(assets.Concat(), assets.Uglify())
  .MustWrite("generated/assets/js")
  .DependsOn(jqueryBundle)
```

```go
homepage_css := assets.Dir("assets/css/home")
  .AllFiles().Filter(assets.Concat(), assets.Sass())
  .MustWrite("generated/assets/css")
```
