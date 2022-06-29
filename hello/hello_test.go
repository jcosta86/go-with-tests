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
		got := Hello("Jeff", "English")
		want := "Hello, Jeff"
		assertCorrectMesage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is suplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMesage(t, got, want)
	})
	t.Run("In Portuguese", func(t *testing.T) {
		got := Hello("Jeff", "Portuguese")
		want := "Ola, Jeff"
		assertCorrectMesage(t, got, want)
	})
	t.Run("In french", func(t *testing.T) {
		got := Hello("Jeff", "french")
		want := "Bonjour, Jeff"
		assertCorrectMesage(t, got, want)
	})
}
