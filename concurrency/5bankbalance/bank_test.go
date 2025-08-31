package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_Bank(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)

	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "$289120.00") {
		t.Error("wrong balance return")
	}

}
