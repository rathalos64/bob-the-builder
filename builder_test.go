package bob

import (
	"strings"
	"testing"
)

func TestBob(t *testing.T) {
	expected := "Bob, the Builder"

	builder := Builder{}
	_, err := builder.WriteString("Bob, the Builder")
	if err != nil {
		t.Fatalf("Bob failed by: %v", err)
	}
	observed := builder.String()

	if !strings.EqualFold(observed, expected) {
		t.Fatalf("Bobs are unequal (observed %q, expected %q)", observed, expected)
	}
}
