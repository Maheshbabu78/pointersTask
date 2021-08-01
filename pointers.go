package main

import (
	"fmt"
	"log"
)

func swapNumbers(a, b *int) {
	tempx := *a
	tempy := *b
	*a = tempy
	*b = tempx
}
func main() {
	divideByZero()
	x := int(1)
	y := int(2)
	fmt.Println("initial value of x:", x)
	fmt.Println("initial value of y:", y)
	swapNumbers(&x, &y)
	fmt.Println(&x)
	fmt.Println(&y)
	fmt.Println("deref", *&x)
	fmt.Println("deref", *&y)
	fmt.Println("x value is:", x)
	fmt.Println("y value is:", y)

}
func divideByZero() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)

		}
	}()
	fmt.Println("the output is ", divide(10, 0))
}

func divide(num1, num2 int) int {
	return num1 / num2
}
