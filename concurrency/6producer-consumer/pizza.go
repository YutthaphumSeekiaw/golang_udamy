package main

import (
	"fmt"
	"math/rand"
	"time"

	color "github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}
type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func main() {
	// seed the random number generate
	rand.Seed(time.Now().UnixNano())

	//print out a message
	color.Cyan("The Pizzaria is open for business")
	color.Cyan("---------------------------------")

	pizzJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// crate a producer
	go pizzeria(pizzJob)

	//run the producer in the background
	for i := range pizzJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery!", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer is really mad!")
			}
		} else {
			color.Cyan("Done making pizzas ....")
			err := pizzJob.Close()
			if err != nil {
				color.Red("*** Error closing channel!", err)
			}
		}
	}

	//create and run consumer

	//print out end message
	color.Cyan("---------------------------------")
	color.Cyan("Done for the day.")

	color.Cyan("We made %d pizzas, but failed to make %d, with %d attempts in total. ", pizzasMade, pizzasFailed, total)

	switch {
	case pizzasFailed > 9:
		color.Red("It was an awful day ....")
	case pizzasFailed >= 6:
		color.Red("It was not a very good day ...")
	case pizzasFailed >= 4:
		color.Yellow("It was ok good day ...")
	case pizzasFailed >= 2:
		color.Yellow("It was pretty good day ...")
	default:
		color.Green("It was great day ...")
	}
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Recieved order #%d! \n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("Making pizza #%d. It will take %d secounds .... \n", pizzaNumber, delay)

		//delay for a bit
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredient for pizza #%d!", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza #%d!", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza Order #%d is ready!", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

		return &p

	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {
	//keep track of which pizza we are making
	var i = 0

	//run forever or until we recieve a quit notification
	//try make pizzas

	for {
		currentPizza := makePizza(i)
		// try to make a pizza
		//decision
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			//we tried to make pizza  (we can someting to the data chanel)
			case pizzaMaker.data <- *currentPizza:
			case quitChab := <-pizzaMaker.quit:
				close(pizzaMaker.data)
				close(quitChab)
				return
			}
		}
	}
}
