package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Berkeli")
	want := "Hello, Berkeli!"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}