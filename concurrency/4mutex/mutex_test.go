package main

import "testing"

func Test_updateMessage(t *testing.T) {
	msg = "Hello, World"

	wg.Add(2)
	//go updateMessage("x")
	go updateMessage("Goodbye, curel world")
	wg.Wait()

	if msg != "Goodbye, curel world" {
		t.Error("incorrect value")
	}
}
