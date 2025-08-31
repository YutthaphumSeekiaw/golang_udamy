package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amout  int64
}

func main() {
	// variable for bank balance
	var bankBalance int64
	var balance sync.Mutex

	//print out starstring value
	fmt.Printf("Initial account balance: $%d.00", bankBalance)
	fmt.Println()

	//define weekly revenue
	incomes := []Income{
		{Source: "Main job", Amout: 5000},
		{Source: "Gifts", Amout: 10},
		{Source: "Part time job", Amout: 50},
		{Source: "Investment", Amout: 500},
	}

	wg.Add(len(incomes))
	//loop through 52 weeek and print out how mush is made ;  keep runing local
	for i, income := range incomes {

		go func(i int, income Income) {
			defer wg.Done()

			for week := 1; week <= 52; week++ {
				balance.Lock()
				tmp := bankBalance
				tmp += income.Amout
				bankBalance = tmp
				balance.Unlock()

				fmt.Printf("On week %d , you earned $%d.00 from %s \n", week, income.Amout, income.Source)
			}
		}(i, income)
	}
	wg.Wait()

	//print out final balance
	fmt.Printf("Final bank balance : $%d.00", bankBalance)
	fmt.Println()
}
