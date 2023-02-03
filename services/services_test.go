package services

import (
	"testing"
)

func TestRandomOddNumber(t *testing.T) {

	max := 2
	min := 1
	want := 1
	got := randomOddNumber(min, max)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

func TestRandomEvenNumber(t *testing.T) {

	max := 3
	min := 1
	want := 2
	got := randomEvenNumber(min, max)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

func TestSum(t *testing.T) {

	arr := []int{1, 2, 10, 20}
	want := 33
	got := sum(arr)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}
