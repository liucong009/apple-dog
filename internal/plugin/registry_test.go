package plugin

import (
	"context"
	"testing"

	"github.com/liucong009/apple-dog/internal/island"
)

type stubPlugin struct {
	desc Descriptor
	snap island.Presentation
}

func (s stubPlugin) Descriptor() Descriptor { return s.desc }
func (s stubPlugin) Start(context.Context, HostCallbacks) error { return nil }
func (s stubPlugin) Stop(context.Context) error                 { return nil }
func (s stubPlugin) Snapshot() island.Presentation              { return s.snap }

func TestRegistry_duplicate(t *testing.T) {
	r := NewRegistry()
	p := stubPlugin{desc: Descriptor{ID: "a", Priority: 1}, snap: island.Presentation{PluginID: "a"}}
	if err := r.Register(p); err != nil {
		t.Fatal(err)
	}
	if err := r.Register(p); err == nil {
		t.Fatal("expected duplicate error")
	}
}

func TestHost_primaryPriority(t *testing.T) {
	r := NewRegistry()
	low := stubPlugin{
		desc: Descriptor{ID: "low", Priority: 5},
		snap: island.Presentation{PluginID: "low", Compact: island.CompactLayout{Leading: island.Slot{Text: "L"}}},
	}
	high := stubPlugin{
		desc: Descriptor{ID: "high", Priority: 1},
		snap: island.Presentation{PluginID: "high", Compact: island.CompactLayout{Leading: island.Slot{Text: "H"}}},
	}
	_ = r.Register(low)
	_ = r.Register(high)
	host := NewHost(r, nil)
	ctx := context.Background()
	_ = host.Activate(ctx, "low")
	_ = host.Activate(ctx, "high")
	primary, err := host.PrimaryPresentation()
	if err != nil {
		t.Fatal(err)
	}
	if primary.PluginID != "high" {
		t.Fatalf("got %q want high", primary.PluginID)
	}
}
