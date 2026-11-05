package plugin

import (
	"context"
	"fmt"
	"sort"

	"github.com/liucong009/apple-dog/internal/island"
)

// Host orchestrates plugins and produces the island presentation for preview or (future) iOS bridge.
type Host struct {
	Registry *Registry
	callbacks HostCallbacks
}

type hostCallbacks struct {
	log func(string)
}

func (h hostCallbacks) RequestRefresh(reason string) {
	if h.log != nil {
		h.log("refresh: " + reason)
	}
}

func (h hostCallbacks) LogInfo(msg string) {
	if h.log != nil {
		h.log(msg)
	}
}

// NewHost wires a registry with default callbacks (optional logger).
func NewHost(reg *Registry, log func(string)) *Host {
	return &Host{
		Registry: reg,
		callbacks: hostCallbacks{log: log},
	}
}

// Activate enables a plugin through the registry.
func (h *Host) Activate(ctx context.Context, id string) error {
	return h.Registry.Activate(ctx, id, h.callbacks)
}

// Deactivate stops a plugin.
func (h *Host) Deactivate(ctx context.Context, id string) error {
	return h.Registry.Deactivate(ctx, id)
}

// PrimaryPresentation picks the highest-priority active plugin snapshot.
func (h *Host) PrimaryPresentation() (island.Presentation, error) {
	active := h.Registry.ActivePlugins()
	if len(active) == 0 {
		return island.Presentation{}, fmt.Errorf("no active plugins")
	}
	sort.Slice(active, func(i, j int) bool {
		return active[i].Descriptor().Priority < active[j].Descriptor().Priority
	})
	for _, p := range active {
		snap := p.Snapshot()
		if err := island.Validate(snap); err != nil {
			continue
		}
		return snap, nil
	}
	return island.Presentation{}, fmt.Errorf("no valid presentation from active plugins")
}

// PresentationFor returns snapshot for a specific plugin id (need not be active; for preview).
func (h *Host) PresentationFor(id string) (island.Presentation, error) {
	p, ok := h.Registry.Get(id)
	if !ok {
		return island.Presentation{}, fmt.Errorf("unknown plugin %q", id)
	}
	snap := p.Snapshot()
	if err := island.Validate(snap); err != nil {
		return island.Presentation{}, err
	}
	return snap, nil
}

// MinimalGlyphs collects minimal glyphs from other active plugins (draft merge).
func (h *Host) MinimalGlyphs(primaryID string) []string {
	active := h.Registry.ActivePlugins()
	var glyphs []string
	for _, p := range active {
		if p.Descriptor().ID == primaryID {
			continue
		}
		g := p.Snapshot().Minimal.Glyph
		if g != "" {
			glyphs = append(glyphs, g)
		}
	}
	return glyphs
}
