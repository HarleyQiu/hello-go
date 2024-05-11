package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeatd := Repeat("a")
	expected := "aaaaa"
	if repeatd != expected {
		t.Errorf("expected %q repeat %q", expected, repeatd)
	}
}
