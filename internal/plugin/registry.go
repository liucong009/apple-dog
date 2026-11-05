package plugin

import (
	"context"
	"fmt"
	"sync"
)

// Lifecycle state of a registered plugin instance.
type Lifecycle string

const (
	LifecycleRegistered Lifecycle = "registered"
	LifecycleActive     Lifecycle = "active"
	LifecycleStopped    Lifecycle = "stopped"
)

type entry struct {
	plugin Plugin
	state  Lifecycle
}

// Registry holds plugins and their lifecycle.
type Registry struct {
	mu      sync.RWMutex
	entries map[string]*entry
}

// NewRegistry creates an empty registry.
func NewRegistry() *Registry {
	return &Registry{entries: make(map[string]*entry)}
}

// Register adds a plugin; id must be unique.
func (r *Registry) Register(p Plugin) error {
	if p == nil {
		return fmt.Errorf("register: nil plugin")
	}
	id := p.Descriptor().ID
	if id == "" {
		return fmt.Errorf("register: empty id")
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.entries[id]; ok {
		return fmt.Errorf("register: duplicate id %q", id)
	}
	r.entries[id] = &entry{plugin: p, state: LifecycleRegistered}
	return nil
}

// List returns descriptors in arbitrary order.
func (r *Registry) List() []Descriptor {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]Descriptor, 0, len(r.entries))
	for _, e := range r.entries {
		out = append(out, e.plugin.Descriptor())
	}
	return out
}

// State returns lifecycle state for id.
func (r *Registry) State(id string) (Lifecycle, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	e, ok := r.entries[id]
	if !ok {
		return "", false
	}
	return e.state, true
}

// Activate starts plugin by id.
func (r *Registry) Activate(ctx context.Context, id string, host HostCallbacks) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	e, ok := r.entries[id]
	if !ok {
		return fmt.Errorf("activate: unknown plugin %q", id)
	}
	if e.state == LifecycleActive {
		return nil
	}
	if err := e.plugin.Start(ctx, host); err != nil {
		return err
	}
	e.state = LifecycleActive
	return nil
}

// Deactivate stops plugin by id.
func (r *Registry) Deactivate(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	e, ok := r.entries[id]
	if !ok {
		return fmt.Errorf("deactivate: unknown plugin %q", id)
	}
	if e.state != LifecycleActive {
		e.state = LifecycleStopped
		return nil
	}
	if err := e.plugin.Stop(ctx); err != nil {
		return err
	}
	e.state = LifecycleStopped
	return nil
}

// ActivePlugins returns plugins in active state.
func (r *Registry) ActivePlugins() []Plugin {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var out []Plugin
	for _, e := range r.entries {
		if e.state == LifecycleActive {
			out = append(out, e.plugin)
		}
	}
	return out
}

// Get returns plugin by id.
func (r *Registry) Get(id string) (Plugin, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	e, ok := r.entries[id]
	if !ok {
		return nil, false
	}
	return e.plugin, true
}
