package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("w1 w2 w3 w4 w5 w6\n")
	exp := 6
	got := count(b)

	if got != exp {
		t.Errorf("Expected %d, got %d.\n", exp, got)
	}
}
