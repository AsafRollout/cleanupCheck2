package fm

import (
	"testing"

	"github.com/rollout/rox-go/v6/core/model"
)

// stubFlag is a minimal model.Flag used to swap the real container field in
// tests. Embedding model.Flag satisfies the interface without implementing
// every method; the test never evaluates the flag through it.
type stubFlag struct {
	model.Flag
}

// TestEnableNewFeature_Swap exercises the reference shape the flag-cleanup
// agent historically missed: the flag is only ever *referenced* as a
// container field (save/restore + assignment for a mock swap), never through
// an evaluation call like `.IsEnabled(...)`. None of the evaluation grep
// seeds match these lines — only the broad `\.EnableNewFeature\b` seed does.
// Removing the container declaration in flags.go orphans these references and
// breaks the build, so this file MUST appear in planned_edits[] alongside the
// declaration removal.
func TestEnableNewFeature_Swap(t *testing.T) {
	original := Flags.EnableNewFeature
	Flags.EnableNewFeature = &stubFlag{}
	defer func() {
		Flags.EnableNewFeature = original
	}()

	if Flags.EnableNewFeature == original {
		t.Fatal("expected the swapped stub flag, got the original")
	}
}
