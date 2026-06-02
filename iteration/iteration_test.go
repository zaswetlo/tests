package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated, called := Repeat("a")
	expected := "aaaaa"
	calledExpected := 5

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}

	if called != calledExpected {
		t.Errorf("expected %d but got %d", calledExpected, called)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}
