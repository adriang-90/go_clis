package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("w1 w2 w3 w4 w5 w6\n")
	exp := 6
	got := count(b, false, false)

	if got != exp {
		t.Errorf("Expected %d, got %d.\n", exp, got)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("w1 w2 \nline2 \nline3 w3")
	exp := 3
	got := count(b, true, false)

	if got != exp {
		t.Errorf("Expected %d, got %d.\n", exp, got)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("dfdfgdfgdfgdfgdfgdfgdfg sdf sdf sdf sdf sdfsdf sdf sdf sdf sdf sdf")
	exp := 66
	got := count(b, false, true)
	if got != exp {
		t.Errorf("Expected %d, got %d.\n", exp, got)
	}
}
