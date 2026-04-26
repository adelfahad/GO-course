package main

import (
	"fmt"

	"adelf.com/manger/bank"
	"adelf.com/manger/calculator"
	"github.com/Pallinder/go-randomdata"
)

func main() {
	fmt.Println("	wecome to manager system")
	fmt.Printf("customers service email: %v\n", randomdata.Email())
	for managerRes() {
	}
}

func managerRes() bool {
	fmt.Print(`
1-bank
2-profit calculator
3-investment calculator
4-exit
`)
	var choice int8
	fmt.Scan(&choice)
	switch choice {
	case 1:
		bank.Bank()
	case 2:
		calculator.ProfitCalculator()
	case 3:
		calculator.InvestmentCalc()
	case 4:
		return false
	default:
		fmt.Println("invalid choice")
	}
	return true
}
