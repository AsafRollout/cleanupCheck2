// Package fmsync is a local stand-in for the CloudBees wrapper
// `github.com/calculi-corp/fmsync-api/pkg/fmsync`, which every
// calculi-corp Go service uses to register its feature-flag container
// instead of calling rox-go directly.
//
// It exists so the fixture can compile on public CI (the real wrapper is
// a private module) while faithfully reproducing the call shape the skill
// must recognize: `fmsync.NewFeatureFlags("<namespace>", ctx, container)`
// registers `container` under `<namespace>` — the first string arg is the
// Rox namespace. The flag's fully-qualified key is therefore
// `<namespace>.<FieldName>`, but source only ever references the bare
// field name.
package fmsync

import (
	"io"

	"github.com/rollout/rox-go/v6/server"
)

// NewFeatureFlags registers flags under svcName (the Rox namespace) and
// returns a closer, mirroring the real wrapper's signature.
func NewFeatureFlags[T any](svcName string, globalCtx map[string]any, flags *T) io.Closer {
	rox := server.NewRox()
	rox.Register(svcName, flags)
	return &closer{}
}

type closer struct{}

func (c *closer) Close() error { return nil }
