package obol

import _ "embed"

//go:embed native/darwin-arm64/libobol_ffi.dylib
var embeddedLib []byte

const embeddedExt = "dylib"
