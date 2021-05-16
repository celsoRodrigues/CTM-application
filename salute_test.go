package main

import "testing"

func TestSalute(t *testing.T) {
	expected := "Hello World"

	if got := Salute(); got != expected {

		t.Errorf("expected %s, got %s", expected, got)

	}
}
