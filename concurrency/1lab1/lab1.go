package main

import (
	"fmt"
	"time"
)

func main() {
	go greet("Hello, World!")
	//time.Sleep(2 * time.Second)
	go showGreet("Hello, greet!")

}

func greet(s string) {
	fmt.Println(s)
}

func showGreet(s string) {
	time.Sleep(1 * time.Second)
	greet(s)
}
