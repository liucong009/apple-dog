package render

import (
	"strings"
	"testing"

	"github.com/liucong009/apple-dog/internal/island"
)

func TestCompactWireframe_containsSlots(t *testing.T) {
	out := CompactWireframe(island.CompactLayout{
		Leading:  island.Slot{Icon: "☀", Text: "24°"},
		Trailing: island.Slot{Text: "晴"},
	})
	if !strings.Contains(out, "24°") || !strings.Contains(out, "晴") {
		t.Fatalf("unexpected output:\n%s", out)
	}
}
