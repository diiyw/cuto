package main

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	if err := Generate(); err != nil {
		t.Fatal(err)
	}
}
