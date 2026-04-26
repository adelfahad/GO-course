package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var revenue, expenses, tax_rate float64
	var earning_before_tax, profit, ratio float64
	var revErr, expErr, taxErr error
	revenue, revErr = userInput("revenue:")
	expenses, expErr = userInput("expenses:")
	tax_rate, taxErr = userInput("tax rate:")
	err := errs(revErr, expErr, taxErr)
	if err != nil {
		return
	}
	earning_before_tax, profit, ratio = calc(revenue, expenses, tax_rate)
	printresults(earning_before_tax, profit, ratio)
}
func calc(revenue, expenses, tax_rate float64) (float64, float64, float64) {
	earning_before_tax := revenue - expenses
	profit := earning_before_tax * (1 - tax_rate/100)
	ratio := earning_before_tax / profit
	result := fmt.Sprintf("earning before tax: %.2f $\nprofit: %.2f $\nratio: %.2f", earning_before_tax, profit, ratio)
	os.WriteFile("result.txt", []byte(result), 0644)
	return earning_before_tax, profit, ratio
}

func userInput(output string) (float64, error) {
	var input float64
	print(output)
	fmt.Scan(&input)
	var err error
	if input < 0 {
		err = errors.New("you cant enter negative number")
	} else if input == 0 {
		err = errors.New("you cant enter 0")
	}
	return input, err
}
func printresults(earning_before_tax, profit, ratio float64) {
	fmt.Printf("earning before tax: %.2f $\n", earning_before_tax)
	fmt.Printf("profit: %.2f $\n", profit)
	var SRatio = fmt.Sprintf("ratio: %.2f", ratio)
	fmt.Print(SRatio)
}

func errs(revErr, expErr, taxErr error) error {
	if revErr != nil {
		fmt.Printf("error revenue: %s \n", revErr)
		return revErr
	} else if expErr != nil {
		fmt.Printf("error expenses: %s \n", expErr)
		return expErr
	} else if taxErr != nil {
		fmt.Printf("error tax rate: %s \n", taxErr)
		return taxErr
	}
	return nil
}
