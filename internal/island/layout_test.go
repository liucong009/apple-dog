package island

import "testing"

func TestValidate_ok(t *testing.T) {
	p := Presentation{
		PluginID: "test",
		Compact: CompactLayout{
			Leading:  Slot{Text: "24°"},
			Trailing: Slot{Text: "晴"},
		},
	}
	if err := Validate(p); err != nil {
		t.Fatal(err)
	}
}

func TestValidate_leadingTooLong(t *testing.T) {
	p := Presentation{
		PluginID: "test",
		Compact: CompactLayout{
			Leading: Slot{Text: "123456789"},
		},
	}
	if err := Validate(p); err == nil {
		t.Fatal("expected error")
	}
}
