package greeting

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGreeting(t *testing.T) {
	name := os.Getenv("NAME")
	got := Greeting(name)
	want := "Hello, Gopher!"
	assert.Equal(t, want, got, "Greeting('Gopher') = %q, want %q", got, want)
}
