package greet

import "fmt"

// Message returns a greeting for name. Empty name uses a generic greeting.
func Message(name string) string {
	if name == "" {
		return "Hello from apple-dog!"
	}
	return fmt.Sprintf("Hello, %s!", name)
}
