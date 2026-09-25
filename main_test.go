package main

import "testing"

func TestGreeting(t *testing.T) {
	if got := greeting("Ada"); got != "Hello, Ada!" {
		t.Fatalf("greeting(%q) = %q", "Ada", got)
	}
}
