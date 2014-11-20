assets
======

Assets management for golang web applications

```go
core := assets.Dir("assets")
  .Files("core.js",
         "util.js",
         "models.js",
         "network.js")
  .Filter(assets.Concat(), assets.Uglify())
  .Write("generated/assets")
  .DependsOn(jqueryBundle)
```
