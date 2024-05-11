package iteration

import "testing"

func TestRepeat(t *testing.T) {

	assertCorrectRepetition := func(t *testing.T, expected, repeatd string) {
		t.Helper()
		if repeatd != expected {
			t.Errorf("expected %q repeat %q", expected, repeatd)
		}
	}

	t.Run("repeat character 'a' five times", func(t *testing.T) {
		repeatd := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectRepetition(t, expected, repeatd)
	})

	t.Run("repeat character 'b' three times", func(t *testing.T) {
		repeatd := Repeat("b", 3)
		expected := "bbb"
		assertCorrectRepetition(t, expected, repeatd)
	})

	t.Run("repeat character 'x' zero times", func(t *testing.T) {
		repeatd := Repeat("x", 0)
		expected := ""
		assertCorrectRepetition(t, expected, repeatd)
	})

}
