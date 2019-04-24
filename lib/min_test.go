package lib

import (
	"testing"
)

func TestMin(t *testing.T) {
	got := Min(1, 10)
	want := 1

	if got != want {
		t.Errorf("Abs(-1) = %d; want %d", got, want)
	}
}
