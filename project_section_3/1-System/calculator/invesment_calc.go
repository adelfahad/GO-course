package calculator

import (
	"fmt"
	"math"
)

func InvestmentCalc() {
	const infltionRate = 3.5
	var balance, investmentTime, expectedReturnRate float64 = 2300, 15, 6.5

	fmt.Print("balance:")
	fmt.Scan(&balance)
	fmt.Print("investmentTime:")
	fmt.Scan(&investmentTime)
	fmt.Print("expectedReturnRate:")
	fmt.Scan(&expectedReturnRate)

	value := balance * math.Pow((1+expectedReturnRate/100.0), investmentTime)
	realValue := value / math.Pow(1+infltionRate/100, investmentTime)

	fmt.Println(value)
	fmt.Println(realValue)
}
