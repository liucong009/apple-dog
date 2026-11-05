package weather

import (
	"context"
	"time"

	"github.com/liucong009/apple-dog/internal/island"
	"github.com/liucong009/apple-dog/internal/plugin"
)

// Plugin is a draft weather snapshot for island preview.
type Plugin struct {
	city string
}

// New returns the weather stub plugin.
func New() *Plugin {
	return &Plugin{city: "海淀"}
}

func (p *Plugin) Descriptor() plugin.Descriptor {
	return plugin.Descriptor{
		ID:              "weather",
		DisplayName:     "天气",
		Priority:        10,
		MaxDurationHint: "4h refresh-driven",
	}
}

func (p *Plugin) Start(context.Context, plugin.HostCallbacks) error { return nil }
func (p *Plugin) Stop(context.Context) error                         { return nil }

func (p *Plugin) Snapshot() island.Presentation {
	now := time.Now().Format("15:04")
	return island.Presentation{
		PluginID: "weather",
		Compact: island.CompactLayout{
			Leading:  island.Slot{Icon: "☀", Text: "24°"},
			Trailing: island.Slot{Text: p.city + " 晴"},
		},
		Minimal: island.MinimalLayout{Glyph: "☀"},
		Expanded: island.ExpandedLayout{
			Title:    "北京",
			Subtitle: "晴",
			Badge:    "天气",
			Regions: []island.Region{
				{
					Kind: island.RegionCenter,
					Lines: []string{
						"24°  体感 26°",
						"18:00 22° · 21:00 20° · 00:00 18°",
						"更新 " + now + " (stub)",
					},
				},
				{
					Kind: island.RegionBottom,
					Actions: []island.Action{
						{ID: "refresh", Label: "刷新"},
						{ID: "open", Label: "详情", Primary: true},
					},
				},
			},
		},
		Companion: island.CompanionLayout{
			TabTitle: "天气插件",
			Sections: []island.CompanionSection{
				{Title: "地点", Lines: []string{"默认：北京 · 海淀（占位）"}},
				{Title: "刷新", Lines: []string{"每 30 分钟（未实现）"}},
			},
			FooterNote: "Live Activity / WeatherKit：需 iOS 原生实现",
		},
	}
}
