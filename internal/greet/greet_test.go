package greet

import "testing"

func TestMessage(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", "Hello from apple-dog!"},
		{"named", "world", "Hello, world!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Message(tt.in); got != tt.want {
				t.Fatalf("Message(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
