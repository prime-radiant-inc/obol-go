# obol-go

[![Go Reference](https://pkg.go.dev/badge/github.com/prime-radiant-inc/obol-go.svg)](https://pkg.go.dev/github.com/prime-radiant-inc/obol-go)

Go binding for [**obol**](https://github.com/prime-radiant-inc/obol) — a library that reads agent
transcripts and estimates their USD cost. Dialects: Claude Code, Codex, Pi, Gemini, opencode,
Copilot, Kimi, and obol's own session logs.

```go
import "github.com/prime-radiant-inc/obol-go"
```

```bash
go get github.com/prime-radiant-inc/obol-go
```

The binding has **no cgo** and needs no C toolchain (`CGO_ENABLED=0` works). The platform's native
library is embedded in the module, so a plain `go get` is all a consumer needs on macOS
(arm64/x64) and Linux (x64/arm64).

## Usage

```go
import "github.com/prime-radiant-inc/obol-go"

obol.Version() // e.g. "0.4.1" — the Rust core version

// Dialect is required; pass one of the supported identifiers.
est, err := obol.EstimatePath("transcript.jsonl", "claude")
if err != nil {
	// *obol.ObolError carries .Code, .Kind, and .Message from the FFI error envelope.
	log.Fatal(err)
}
fmt.Println(est.TotalUSD, est.PricingAsOf)
for _, m := range est.PerModel {
	fmt.Println(m.Model, m.Provider, m.SubtotalUSD)
}

// Refresh the on-disk pricing snapshot (network call).
report, err := obol.Refresh("2026-06-12")
```

Dialect identifiers: `claude`, `codex`, `pi`, `gemini`, `opencode`, `copilot`, `kimi`, `obol`. An
empty, unknown, or invalid dialect returns an `*ObolError` (`Kind == "InvalidArgument"` or
`"UnknownDialect"`) — there is no auto-detection.

### Pricing tables must exist

`EstimatePath` reads a pricing snapshot from disk. Either run `obol refresh` (the CLI) / call
`obol.Refresh(...)`, or point `OBOL_PRICING_DIR` at a directory containing `current.json`. With no
snapshot the call returns an `*ObolError` with `Kind == "PricingTablesMissing"` (code 1).

> On Linux with `CGO_ENABLED=0`, a *runtime* `os.Setenv("OBOL_PRICING_DIR", …)` does **not** reach
> the dlopen'd native library — Go makes raw syscalls and never links libc, so the library's
> `getenv` won't see it. Set the variable **before** the process starts. Inherited environment is
> fine everywhere.

### Pointing at a specific library

Set `OBOL_LIB` to an explicit path to override the embedded library. On macOS under a hardened
runtime with library validation, an unsigned extracted dylib may be rejected — point `OBOL_LIB` at
a signed copy in that case.

## How it works

The binding loads obol's native library at runtime via
[purego](https://github.com/ebitengine/purego) (`dlopen`) and re-types the JSON the Rust core
returns into idiomatic Go structs. The Rust core stays the single source of truth for all
accounting; this package only marshals C strings and unmarshals JSON. In the published module the
platform library is embedded and extracted to a content-hashed dir under `os.UserCacheDir()` (or
the temp dir) on first use, then `dlopen`'d.

## This is a generated repository

The source of truth lives in the main repo under
[`bindings/go/`](https://github.com/prime-radiant-inc/obol/tree/main/bindings/go). This repository
is **populated automatically by obol's release workflow**: each `vX.Y.Z` release assembles the Go
source together with the prebuilt native libraries and tags a matching release here, so `go get`
resolves a self-contained module.

- **File issues and PRs upstream:** https://github.com/prime-radiant-inc/obol
- **Do not** hand-edit the generated module files (`*.go`, `go.mod`, `native/`) — they are
  overwritten on the next release.

## License

Apache-2.0. See [LICENSE](LICENSE) and [NOTICE](NOTICE).
