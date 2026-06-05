# obol-go

Go binding for [**obol**](https://github.com/prime-radiant-inc/obol) — a library that
reads agent transcripts (Claude Code, Codex, Pi) and estimates their USD cost.

```go
import "github.com/prime-radiant-inc/obol-go"
```

## This is a generated repository

The source of truth lives in the main repo under
[`bindings/go/`](https://github.com/prime-radiant-inc/obol/tree/main/bindings/go). This
repository is **populated automatically by obol's release workflow**: each `vX.Y.Z` release
assembles the Go source together with the prebuilt native libraries and tags a matching release
here, so `go get` resolves a self-contained module.

- **File issues and PRs upstream:** https://github.com/prime-radiant-inc/obol
- **Do not** hand-edit this repository — changes here are overwritten on the next release.

## Status

Bootstrap. The first tagged release is pending; until a `vX.Y.Z` tag exists here, `go get` has
nothing to resolve. Watch the [main repo](https://github.com/prime-radiant-inc/obol) for the
release that seeds this module.

## How it works (once released)

The binding loads obol's native library at runtime via
[purego](https://github.com/ebitengine/purego) — **no cgo, no C toolchain required**
(`CGO_ENABLED=0` works). The platform's native library is embedded in the module and used
directly, so a plain `go get` is all a consumer needs on macOS (arm64/x64) and Linux (x64/arm64).

## License

Apache-2.0. See [LICENSE](LICENSE) and [NOTICE](NOTICE).
