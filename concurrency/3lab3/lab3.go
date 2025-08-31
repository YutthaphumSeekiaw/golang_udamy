package main

import (
	"fmt"
	"sync"
)

var msg string

func main() {
	var wg sync.WaitGroup
	msg = "hello, univers"

	wg.Add(1)
	go updateMessage("Hello ,univest")
	printMessage()

	wg.Add(1)
	go updateMessage("Hello ,cosmo")
	printMessage()

	wg.Add(1)
	go updateMessage("Hello ,world")
	wg.Wait()
	printMessage()
	wg.Done()
}

func updateMessage(s string) {
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}
