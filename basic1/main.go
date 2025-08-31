package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// var invertment float64
	// expect := 1.05
	// years := 10.0

	// const inflationRate = 0.02

	// fmt.Println("Enter the amount you want to invest:")
	// fmt.Scan(&invertment)

	// rate := math.Pow(1+expect/100, float64(years))
	// fmt.Printf("The rate of return after %d years is: %.2f\n", years, rate)

	// reulult := float64(invertment) * math.Pow(1+expect/100, float64(years))
	// fmt.Printf("After %d years, the investment will be worth: %.2f\n", years, reulult)

	// fetureResult := invertment * math.Pow(1+inflationRate/100, float64(years))
	// fmt.Printf("After %d years, the investment adjusted for inflation will be worth: %.2f\n", years, fetureResult)

	// concatstring := fmt.Sprintf("concat Investment: %.2f", fetureResult)
	// fmt.Println(concatstring)

	revenue, err := getValue("revenue :")
	if err != nil {
		fmt.Println("Error getting revenue:", err)
		// Exit the program or handle the error as needed
	}

	expenses, err := getValue("expenses :")
	if err != nil {
		fmt.Println("Error getting expenses:", err)
		// Exit the program or handle the error as needed
	}
	taxrate, err := getValue("tax rate :")
	if err != nil {
		fmt.Println("Error getting tax rate:", err)
		// Exit the program or handle the error as needed
	}

	profit, tax, edb, ratio := calculateProfit(revenue, expenses, taxrate)
	fmt.Printf("The profit after tax is: %.2f\n", profit)
	fmt.Printf("The tax amount is: %.2f\n", tax)
	fmt.Printf("The earnings before tax is: %.2f\n", edb)
	fmt.Printf("The tax ratio is: %.2f\n", ratio)

	var choice int
	fmt.Print("choiece: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("You chose option 1")
	} else if choice == 2 {
		fmt.Println("You chose option 2")
	} else if choice == 3 {
		fmt.Println("You chose option 3")
	} else {
		fmt.Println("Invalid choice")
	}

	whiteToFile("ทดสอบการเขียนไฟล์")
	data, err := readFile()
	if err != nil {
		fmt.Println("Error reading file:", err)
		panic(err) // or handle the error as needed
		// You can choose to return or exit the program here if needed
		// For example, you could use os.Exit(1) to exit with a non-zero status
		// os.Exit(1)
		//return
	}
	fmt.Println("Data read from file:", data)
	fmt.Println("Program completed successfully.")

	age := 34
	var agePointer *int
	agePointer = &age
	fmt.Println("Age:", age)
	fmt.Println("Age Pointer:", *agePointer)

	getAdultAge(agePointer)
	fmt.Println("Age after adjustment:", age)
	fmt.Println("Age Pointer after adjustment:", *agePointer)

}

func getAdultAge(age *int) {
	*age = *age - 18
}

func getValue(text string) (float64, error) {
	var value float64
	fmt.Println(text)
	fmt.Scan(&value)
	if value < 0 {
		return 0, errors.New("value cannot be negative")
	}

	fmt.Printf("You entered: %.2f\n", value)
	// Here you can add any additional logic if needed, such as writing to a file
	return value, nil
}

func calculateProfit(revenue, expenses, taxRate float64) (float64, float64, float64, float64) {
	edb := revenue - expenses
	tax := edb * (taxRate / 100)
	ratio := tax / edb
	return edb - tax, tax, edb, ratio
}

func whiteToFile(text string) {
	// This function is a placeholder for writing to a file.
	// Implementation would go here if needed.
	os.WriteFile("text.txt", []byte(text), 0644)
	fmt.Println("Text written to file successfully.")
}

func readFile() (string, error) {
	data, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return "", err
	}
	fmt.Println("File read successfully.", string(data))

	text := "30"
	number, err := strconv.ParseFloat(text, 64)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return "", err
	}
	fmt.Println("Converted number:", number)
	return string(data), nil
}
