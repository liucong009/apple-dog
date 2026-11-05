// Package render turns island.Presentation into ASCII wireframes for CLI preview.
package render

import (
	"fmt"
	"strings"

	"github.com/liucong009/apple-dog/internal/island"
)

const frameWidth = 44

// CompactWireframe renders the pill compact layout.
func CompactWireframe(c island.CompactLayout) string {
	left := formatSlot(c.Leading)
	right := formatSlot(c.Trailing)
	inner := fmt.Sprintf("%s ··· pill ··· %s", left, right)
	return box("Compact (Dynamic Island pill)", inner)
}

// ExpandedWireframe renders expanded regions.
func ExpandedWireframe(e island.ExpandedLayout) string {
	var b strings.Builder
	header := e.Title
	if e.Subtitle != "" {
		header = header + " · " + e.Subtitle
	}
	if e.Badge != "" {
		header = padLine(header, e.Badge)
	}
	b.WriteString(line(header))
	b.WriteString(sep())
	for i, r := range e.Regions {
		for _, ln := range r.Lines {
			b.WriteString(line(ln))
		}
		if len(r.Actions) > 0 {
			var labels []string
			for _, a := range r.Actions {
				if a.Primary {
					labels = append(labels, "["+a.Label+"*]")
				} else {
					labels = append(labels, "["+a.Label+"]")
				}
			}
			b.WriteString(line(strings.Join(labels, "  ")))
		}
		if i < len(e.Regions)-1 {
			b.WriteString(sep())
		}
	}
	return box("Expanded", strings.TrimRight(b.String(), "\n"))
}

// CompanionWireframe renders host app companion tab.
func CompanionWireframe(c island.CompanionLayout) string {
	var b strings.Builder
	title := c.TabTitle
	if title == "" {
		title = "Companion"
	}
	b.WriteString(line("Tab: " + title))
	b.WriteString(sep())
	for _, sec := range c.Sections {
		b.WriteString(line("▸ " + sec.Title))
		for _, ln := range sec.Lines {
			b.WriteString(line("  " + ln))
		}
	}
	if c.FooterNote != "" {
		b.WriteString(sep())
		b.WriteString(line(c.FooterNote))
	}
	return box("Companion App", strings.TrimRight(b.String(), "\n"))
}

// FullPreview renders compact + expanded + optional minimal glyphs.
func FullPreview(p island.Presentation, extraMinimal []string) string {
	var parts []string
	parts = append(parts, CompactWireframe(p.Compact))
	if p.Minimal.Glyph != "" || len(extraMinimal) > 0 {
		glyphs := p.Minimal.Glyph
		if len(extraMinimal) > 0 {
			glyphs = strings.Join(append([]string{p.Minimal.Glyph}, extraMinimal...), " ")
		}
		parts = append(parts, box("Minimal", line("glyphs: "+strings.TrimSpace(glyphs))))
	}
	parts = append(parts, ExpandedWireframe(p.Expanded))
	return strings.Join(parts, "\n\n")
}

func formatSlot(s island.Slot) string {
	if s.Icon != "" && s.Text != "" {
		return s.Icon + " " + s.Text
	}
	if s.Icon != "" {
		return s.Icon
	}
	return s.Text
}

func box(title, body string) string {
	top := "┌─ " + title + strings.Repeat("─", max(0, frameWidth-len(title)-4)) + "┐"
	var lines []string
	for _, ln := range strings.Split(body, "\n") {
		lines = append(lines, "│"+padRight(ln, frameWidth)+"│")
	}
	bottom := "└" + strings.Repeat("─", frameWidth+2) + "┘"
	return strings.Join(append([]string{top}, lines...), "\n") + "\n" + bottom
}

func line(s string) string {
	return padRight(s, frameWidth)
}

func sep() string {
	return strings.Repeat("─", frameWidth) + "\n"
}

func padLine(left, right string) string {
	space := frameWidth - len(left) - len(right)
	if space < 1 {
		return left + " " + right
	}
	return left + strings.Repeat(" ", space) + right
}

func padRight(s string, w int) string {
	if len(s) >= w {
		return s[:w]
	}
	return s + strings.Repeat(" ", w-len(s))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
