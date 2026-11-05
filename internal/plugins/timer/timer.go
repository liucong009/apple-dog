package timer

import (
	"context"
	"fmt"
	"time"

	"github.com/liucong009/apple-dog/internal/island"
	"github.com/liucong009/apple-dog/internal/plugin"
)

// Plugin is a draft pomodoro-style timer stub.
type Plugin struct {
	deadline time.Time
}

// New returns a timer stub with a fixed offset for deterministic previews.
func New() *Plugin {
	return &Plugin{deadline: time.Now().Add(12*time.Minute + 34*time.Second)}
}

func (p *Plugin) Descriptor() plugin.Descriptor {
	return plugin.Descriptor{
		ID:              "timer",
		DisplayName:     "计时",
		Priority:        20,
		MaxDurationHint: "until timer ends",
	}
}

func (p *Plugin) Start(ctx context.Context, host plugin.HostCallbacks) error {
	host.LogInfo("timer started (stub)")
	return nil
}

func (p *Plugin) Stop(context.Context) error { return nil }

func (p *Plugin) Snapshot() island.Presentation {
	remaining := time.Until(p.deadline).Round(time.Second)
	if remaining < 0 {
		remaining = 0
	}
	mmss := formatMMSS(remaining)
	return island.Presentation{
		PluginID: "timer",
		Compact: island.CompactLayout{
			Leading:  island.Slot{Icon: "⏱", Text: "番茄"},
			Trailing: island.Slot{Text: mmss},
		},
		Minimal: island.MinimalLayout{Glyph: "⏱"},
		Expanded: island.ExpandedLayout{
			Title:    "番茄钟",
			Subtitle: mmss + " 剩余",
			Badge:    "计时",
			Regions: []island.Region{
				{
					Kind:  island.RegionCenter,
					Lines: []string{"阶段：专注（占位）", "下一休息 5 分钟"},
				},
				{
					Kind: island.RegionBottom,
					Actions: []island.Action{
						{ID: "pause", Label: "暂停"},
						{ID: "done", Label: "完成", Primary: true},
					},
				},
			},
		},
		Companion: island.CompanionLayout{
			TabTitle: "计时插件",
			Sections: []island.CompanionSection{
				{Title: "预设", Lines: []string{"25 / 5 番茄（占位）"}},
			},
			FooterNote: "倒计时驱动 Activity.update：需 iOS 原生实现",
		},
	}
}

func formatMMSS(d time.Duration) string {
	m := int(d.Minutes())
	s := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d", m, s)
}
