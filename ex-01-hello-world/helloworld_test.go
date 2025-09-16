package main

import (
	"bytes"
	"os"
	"testing"
)

func TestMainOutput(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe() failed: %v", err)
	}
	old := os.Stdout
	os.Stdout = w
	defer func() { os.Stdout = old }()

	main()
	w.Close()

	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	if err != nil {
		t.Fatalf("ReadFrom pipe failed: %v", err)
	}
	got := buf.String()
	want := "Hello World from KT\n"
	if got != want {
		t.Errorf("main() output = %q, want %q", got, want)
	}
}
