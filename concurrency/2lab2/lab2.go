package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	word := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
	}

	wg.Add(len(word))

	for i, x := range word {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	wg.Wait()

	wg.Add(1)

	printSomething("This is the secound thing to be printed!", &wg)
}

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(s)
}
