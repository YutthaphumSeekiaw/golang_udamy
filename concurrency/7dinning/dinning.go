package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	name      string
	leftFork  int
	rightFork int
}

var philosophers = []Philosopher{
	{name: "Plato", leftFork: 0, rightFork: 1},
	{name: "Socrates", leftFork: 1, rightFork: 2},
	{name: "Aristotle", leftFork: 2, rightFork: 3},
	{name: "Kant", leftFork: 3, rightFork: 4},
	{name: "Nietzsche", leftFork: 4, rightFork: 0},
}

var hanger = 3
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

func main() {
	// print out welcome message
	fmt.Println("Welcome to the Dining Philosophers Problem Simulation!")
	fmt.Println("------------------------------------------------")
	fmt.Println("The Table is empty. Let's invite some philosophers to dine.")

	//start the meal
	dine()
	// print out goodbye message
	fmt.Println("The table id emty. The philosophers have left. Goodbye!")
}

func dine() {
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	//forks is map all 5 forks
	forks := make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	//start each meal
	for i := 0; i < len(philosophers); i++ {
		//fire off a goroutine for each philosopher
		go diningProblem(philosophers[i], wg, forks, seated)
	}
	wg.Wait()
}

func diningProblem(p Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%s is seated at the table.\n", p.name)
	seated.Done()
	seated.Wait()

	//eat times
	for i := hanger; i > 0; i-- {
		forks[p.leftFork].Lock()
		fmt.Printf("\t%s picks up left fork .\n", p.name)
		forks[p.rightFork].Lock()
		fmt.Printf("\t%s picks up right fork .\n", p.name)

		fmt.Printf("\t%s is eating. %d\n", p.name, hanger-i+1)
		time.Sleep(eatTime)

		fmt.Printf(("\t%s is Think .\n"), p.name)
		time.Sleep(thinkTime)

		forks[p.rightFork].Unlock()
		fmt.Printf("\t%s puts down right fork .\n", p.name)
		forks[p.leftFork].Unlock()
		fmt.Printf("\t%s puts down left fork .\n", p.name)
	}

	fmt.Printf("%s is satisfied and leaves the table.\n", p.name)

}
