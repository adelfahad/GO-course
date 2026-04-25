package main

import (
	"fmt"
)

var balance float64

func main() {
	fmt.Println("Welcome to adelf bank app")
	i := true
	for i {
		i = selection()
	}
	fmt.Println("thank you for using adelf bank")
}

func selection() bool {
	fmt.Println(`
	wich you what you want to do
	1-check balance 
	2-Deposit amount 
	3-Withdraw money 
	4-exit `)
	var choice int8
	fmt.Scan(&choice)
	return response(choice)
}
func response(choice int8) bool {
	var money float64
	switch choice {
	case 1:
		fmt.Printf("your balance is %.2f$\n", balance)
	case 2:
		fmt.Print("enter deposite amount ")
		fmt.Scan(&money)
		if money > 0 {
			balance += money
		} else {
			fmt.Println("enter a valid number")
		}
	case 3:
		fmt.Print("enter withdrw money ")
		fmt.Scan(&money)
		if money <= balance && money > 0 {
			balance -= money
		} else if money > balance {
			fmt.Println("your balance is low")
		} else {
			fmt.Println("enter a valid number")
		}
	case 4:
		return false
	default:
		fmt.Println("enter a valid number")
	}
	return true
}
