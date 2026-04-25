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
	profit = profit_calc(earning_before_tax, tax_rate)
	ratio = earning_before_tax / profit
	fmt.Printf("earning before tax: %.2f $\n", earning_before_tax)
	fmt.Printf("profit: %.2f $\n", profit)
	var SRatio = fmt.Sprintf("ratio: %.2f", ratio)
	fmt.Print(SRatio)
}
func profit_calc(earning_before_tax, tax_rate float64) float64 {
	return earning_before_tax * (1 - tax_rate/100)
}
