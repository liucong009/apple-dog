package plugin

import (
	"context"

	"github.com/liucong009/apple-dog/internal/island"
)

// Descriptor is plugin metadata for registry and future App Store manifest.
type Descriptor struct {
	ID          string
	DisplayName string
	Priority    int // lower number wins primary election
	// MaxDurationHint documents expected Live Activity lifetime (informational).
	MaxDurationHint string
}

// HostCallbacks lets plugins ask the host to refresh island content (CLI no-op stub).
type HostCallbacks interface {
	RequestRefresh(reason string)
	LogInfo(msg string)
}

// Plugin is the integration surface for Dynamic Island expansion modules.
type Plugin interface {
	Descriptor() Descriptor
	Start(ctx context.Context, host HostCallbacks) error
	Stop(ctx context.Context) error
	Snapshot() island.Presentation
}
