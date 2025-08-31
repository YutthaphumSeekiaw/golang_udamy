package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

type Product struct {
	ID    int
	Name  string
	Price float64
}

func main() {
	// website := map[string]string{
	// 	"web1": "https://example.com",
	// 	"test": "https://test.com",
	// }

	// fmt.Println(website)
	// fmt.Println(website["web1"])

	// delete(website, "test")
	// fmt.Println(website)

	// website["new"] = "https://newsite.com"
	// fmt.Println(website)

	// for key, value := range website {
	// 	fmt.Printf("Key: %s, Value: %s\n", key, value)
	// }

	users := map[string]User{
		"user1": {Name: "Alice", Email: "dfdff"},
		"user2": {Name: "Bob", Email: "sdfsdf"},
	}

	fmt.Println(users)
	fmt.Println(users["user1"])

	products := make(map[int]Product)
	products[101] = Product{ID: 101, Name: "Laptop", Price: 999.99}
	products[102] = Product{ID: 102, Name: "Smartphone", Price: 499.99}

	fmt.Println(products)

	for id, product := range products {
		fmt.Printf("Product ID: %d, Name: %s, Price: %.2f\n", id, product.Name, product.Price)
	}

	numbers := []int{1, 2, 3, 4, 5}

	dNumDouber := transformsNumber(&numbers, doubleNumber)
	fmt.Println(dNumDouber)

	dNumTrirple := transformsNumber(&numbers, trirpleNumber)
	fmt.Println(dNumTrirple)

	//recursive function

	fac := factorial(5)

	fmt.Println(fac)

	//variadic function
	//sum := sumUp([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	sum := sumUp(1, 2, 3, 4, 5, 6, 7, 8, 9, -10)
	fmt.Println(sum)
}

type transformFn func(int) int

// type anotherFn func(int ,[]string , map[string]string) ([]int,string)

func transformsNumber(numbers *[]int, transfromfn transformFn) []int {
	dNumber := []int{}
	for _, number := range *numbers {
		dNumber = append(dNumber, transfromfn(number))
	}
	return dNumber
}

func doubleNumber(number int) int {
	return number * 2
}

func trirpleNumber(number int) int {
	return number * 3
}

// func getTranformerFunction(number *int) transformer{
// 	if (*number % 2) == 0 {
// 		return doubleNumber
// 	}
// 	return double
// }

func factorial(n int) int {
	// res := 1
	// for i := 1; i <= n; i++ {
	// 	res = res * i
	// }
	// return res
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// func sumUp(number []int) int {
func sumUp(number ...int) int {
	sum := 0
	for _, num := range number {
		sum += num // sum = sum + num
	}
	return sum
}
