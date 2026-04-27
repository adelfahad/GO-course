package main

import "fmt"

func main() {
	var x = 5
	var y = 18
	swip(&x, &y)
	fmt.Print(x, y)
}
func swip(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
