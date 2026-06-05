package obol

import _ "embed"

//go:embed native/linux-x64/libobol_ffi.so
var embeddedLib []byte

const embeddedExt = "so"
