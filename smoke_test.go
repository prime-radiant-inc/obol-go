package obol

import "testing"

func TestVersionLoadsEmbedded(t *testing.T) {
	if v := Version(); v == "" {
		t.Fatal("Version() empty — the embedded library failed to load")
	}
}
