package main

import "testing"

func TestSomething(t *testing.T) {
	a := 1
	b := 2
	want := 3
	got := a + b
	if want != got {
		t.Fatalf("wanted %d but got %d", want, got)
	}
}
