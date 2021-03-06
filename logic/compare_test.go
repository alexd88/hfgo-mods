package logic

import "testing"

func TestFirstLarger(t *testing.T) {
	want := 2
	got := Larger(2, 1)
	if got != want {
		t.Errorf("Larger(%d, %d) = %d, want %d", 2, 1, got, want)
	}
}

func TestSecondLarger(t *testing.T) {
	want := 8
	got := Larger(4, 8)
	if got != want {
		t.Errorf("Larger(%d, %d) = %d, want %d", 4, 8, got, want)
	}
}
