package main

import (
	"reflect"
	"testing"
)

func TestMyTime(t *testing.T) {

	myt := myTime()
	expected := "string"

	if got := reflect.TypeOf(myt); got.String() != expected {
		t.Errorf("expected %s, got %s", expected, got)
	}
}
