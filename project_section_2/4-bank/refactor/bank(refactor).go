package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const dataFile = "data.txt"

var balance float64

func main() {
	fmt.Println("Welcome to adelf bank app")
	var err error
	balance, err = read()
	if err != nil {
		fmt.Println(err)
	}
	for selection() {
	}
	fmt.Println("thank you for using adelf bank")
}

func selection() bool {
	fmt.Println(`
	What do you want to do?
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
		write(balance)
	case 2:
		fmt.Print("enter deposite amount ")
		fmt.Scan(&money)
		if money > 0 {
			balance += money
			write(balance)
		} else {
			fmt.Println("enter a valid number")
		}
	case 3:
		fmt.Print("enter withdrw money ")
		fmt.Scan(&money)
		if money <= balance && money > 0 {
			balance -= money
			write(balance)
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

func write(balance float64) {
	dataBalance := fmt.Sprint(balance)
	os.WriteFile(dataFile, []byte(dataBalance), 0644)

}
func read() (float64, error) {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		return 0, errors.New("new account")
	}
	Sdata := string(data)
	balance, err := strconv.ParseFloat(Sdata, 64)
	if err != nil {
		return 0, errors.New("data type error")
	}
	return balance, nil
}
