package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func main() {
	msg = "Hello, world"

	// var mutex sync.Mutex
	wg.Add(2)
	// go updateMessage("Hello, Univese", &mutex)
	// go updateMessage("Hello, Cosmo", &mutex)
	go updateMessage("Hello, Univese")
	go updateMessage("Hello, Cosmo")
	wg.Wait()
	fmt.Println(msg)
}

// func updateMessage(s string, m *sync.Mutex) {
func updateMessage(s string) {
	defer wg.Done()

	//m.Lock()
	msg = s
	//m.Unlock()
}
