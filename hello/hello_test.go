package main

import "testing"

func TestHello(t *testing.T) {
	got := hello("Jeff")
	want := "hello, Jeff"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
