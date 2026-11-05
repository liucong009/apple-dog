// Package plugins registers built-in draft island plugins.
package plugins

import (
	"github.com/liucong009/apple-dog/internal/plugin"
	"github.com/liucong009/apple-dog/internal/plugins/timer"
	"github.com/liucong009/apple-dog/internal/plugins/weather"
)

// RegisterBuiltin adds weather and timer stubs to reg.
func RegisterBuiltin(reg *plugin.Registry) error {
	for _, p := range []plugin.Plugin{weather.New(), timer.New()} {
		if err := reg.Register(p); err != nil {
			return err
		}
	}
	return nil
}
