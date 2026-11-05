package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/liucong009/apple-dog/internal/island"
	"github.com/liucong009/apple-dog/internal/plugin"
	"github.com/liucong009/apple-dog/internal/plugins"
	"github.com/liucong009/apple-dog/internal/render"
)

func main() {
	os.Exit(run())
}

func run() int {
	fs := flag.NewFlagSet("apple-dog", flag.ExitOnError)
	mode := fs.String("mode", "full", "preview: compact | expanded | full | companion")
	pluginID := fs.String("plugin", "weather", "plugin id for preview when -activate is empty")
	list := fs.Bool("list", false, "list registered plugins")
	activate := fs.String("activate", "", "comma-separated ids to activate (primary = best priority)")
	_ = fs.Parse(os.Args[1:])

	reg := plugin.NewRegistry()
	if err := plugins.RegisterBuiltin(reg); err != nil {
		fmt.Fprintf(os.Stderr, "register: %v\n", err)
		return 1
	}

	if *list {
		printPlugins(reg)
		return 0
	}

	host := plugin.NewHost(reg, nil)
	ctx := context.Background()

	if *activate != "" {
		for _, id := range strings.Split(*activate, ",") {
			id = strings.TrimSpace(id)
			if id == "" {
				continue
			}
			if err := host.Activate(ctx, id); err != nil {
				fmt.Fprintf(os.Stderr, "activate %q: %v\n", id, err)
				return 1
			}
		}
	} else if err := host.Activate(ctx, *pluginID); err != nil {
		fmt.Fprintf(os.Stderr, "activate: %v\n", err)
		return 1
	}

	var pres island.Presentation
	var err error
	if *activate != "" {
		pres, err = host.PrimaryPresentation()
	} else {
		pres, err = host.PresentationFor(*pluginID)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "presentation: %v\n", err)
		return 1
	}

	extra := host.MinimalGlyphs(pres.PluginID)

	switch strings.ToLower(*mode) {
	case "compact":
		fmt.Println(render.CompactWireframe(pres.Compact))
	case "expanded":
		fmt.Println(render.ExpandedWireframe(pres.Expanded))
	case "companion":
		fmt.Println(render.CompanionWireframe(pres.Companion))
	case "full":
		fmt.Println(render.FullPreview(pres, extra))
	default:
		fmt.Fprintf(os.Stderr, "unknown mode %q\n", *mode)
		return 2
	}
	return 0
}

func printPlugins(reg *plugin.Registry) {
	descs := reg.List()
	sort.Slice(descs, func(i, j int) bool {
		return descs[i].Priority < descs[j].Priority
	})
	fmt.Println("Registered plugins (priority ascending):")
	for _, d := range descs {
		st, _ := reg.State(d.ID)
		fmt.Printf("  %-8s  %-6s  priority=%d  %s\n", d.ID, d.DisplayName, d.Priority, st)
	}
	fmt.Println("\nDocs: docs/product-concept.md, docs/ui-layout.md, docs/plugin-integration.md")
}
