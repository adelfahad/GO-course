package main

import (
	"fmt"
)

func main() {
	var revenue, expenses, tax_rate float64
	var earning_before_tax, profit, ratio float64
	revenue = userInput("revenue:")
	expenses = userInput("expenses:")
	tax_rate = userInput("tax rate:")
	earning_before_tax, profit, ratio = calc(revenue, expenses, tax_rate)
	fmt.Printf("earning before tax: %.2f $\n", earning_before_tax)
	fmt.Printf("profit: %.2f $\n", profit)
	var SRatio = fmt.Sprintf("ratio: %.2f", ratio)
	fmt.Print(SRatio)
}
func calc(revenue, expenses, tax_rate float64) (float64, float64, float64) {
	earning_before_tax := revenue - expenses
	profit := earning_before_tax * (1 - tax_rate/100)
	ratio := earning_before_tax / profit
	return earning_before_tax, profit, ratio
}

func userInput(output string) float64 {
	var input float64
	print(output)
	fmt.Scan(&input)
	return input
}
