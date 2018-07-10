package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	if true {
		t.Error("True is true")
	}
}

func TestMainFails(t *testing.T) {
	t.Error("Always fails")
}
