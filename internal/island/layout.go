// Package island defines declarative layout models aligned with Dynamic Island
// compact / minimal / expanded regions and companion app screens.
package island

import "fmt"

// RegionKind maps to DynamicIsland expanded regions (Swift side).
type RegionKind string

const (
	RegionTop    RegionKind = "top"
	RegionCenter RegionKind = "center"
	RegionBottom RegionKind = "bottom"
)

// Slot is one compact leading/trailing cell.
type Slot struct {
	Icon string // SF Symbol name placeholder, e.g. "sun.max.fill"
	Text string
}

// CompactLayout is the pill leading/trailing pair.
type CompactLayout struct {
	Leading  Slot
	Trailing Slot
}

// MinimalLayout is shown when multiple activities compete.
type MinimalLayout struct {
	Glyph string
}

// Action is a tappable control in expanded or companion UI (Swift implements).
type Action struct {
	ID    string
	Label string
	Primary bool
}

// Region is one expanded island band.
type Region struct {
	Kind    RegionKind
	Lines   []string
	Actions []Action
}

// ExpandedLayout is the full expanded Dynamic Island content.
type ExpandedLayout struct {
	Title    string
	Subtitle string
	Badge    string
	Regions  []Region
}

// CompanionSection is one block on a companion app tab (wireframe-level).
type CompanionSection struct {
	Title string
	Lines []string
}

// CompanionLayout describes host app screens (not rendered on device island).
type CompanionLayout struct {
	TabTitle   string
	Sections   []CompanionSection
	FooterNote string
}

// Presentation is the full declarative snapshot for one plugin.
type Presentation struct {
	PluginID string
	Compact  CompactLayout
	Minimal  MinimalLayout
	Expanded ExpandedLayout
	Companion CompanionLayout
}

// Validate applies draft length checks for compact slots.
func Validate(p Presentation) error {
	if p.PluginID == "" {
		return fmt.Errorf("presentation: empty plugin id")
	}
	if len([]rune(p.Compact.Leading.Text)) > 8 {
		return fmt.Errorf("compact leading text too long")
	}
	if len([]rune(p.Compact.Trailing.Text)) > 12 {
		return fmt.Errorf("compact trailing text too long")
	}
	return nil
}
