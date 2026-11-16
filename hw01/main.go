package main

import "fmt"

const appName = "Expense Manager"

func main() {
	name := "Lunch"
	amount := 12
	paid := true

	fmt.Println(appName)
	fmt.Printf("Expense: %s | amount: %d € | paid: %v\n", name, amount, paid)
}
