package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMesage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to the people", func(t *testing.T) {
		got := hello("Jeff")
		want := "Hello, Jeff"
		assertCorrectMesage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is suplied", func(t *testing.T) {
		got := hello("")
		want := "Hello, World"
		assertCorrectMesage(t, got, want)
	})
}
