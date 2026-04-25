package main

import (
	"fmt"
)

func main() {
	var revenue, expenses, tax_rate float64
	var earning_before_tax, profit, ratio float64
	fmt.Print("revenue:")
	fmt.Scan(&revenue)
	fmt.Print("expenses:")
	fmt.Scan(&expenses)
	fmt.Print("tax rate:")
	fmt.Scan(&tax_rate)
	earning_before_tax = revenue - expenses
	profit = earning_before_tax * (1 - tax_rate/100)
	ratio = earning_before_tax / profit
	fmt.Print("earning before tax:")
	fmt.Println(earning_before_tax)
	fmt.Print("profit:")
	fmt.Println(profit)
	fmt.Print("ratio:")
	fmt.Println(ratio)
}
