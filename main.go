package main

import "fmt"

func main() {
	fmt.Println(Sum(1, 1))
	fmt.Println(Subtract(1, 1))
	fmt.Println(Multiply(1, 1))
	fmt.Println(Divide(1, 1))
}

func Divide(x float64, y float64) float64 {
	return x / y
}

func Multiply(x float64, y float64) float64 {
	return x * y
}

func Subtract(x float64, y float64) float64 {
	return x - y
}

func Sum(x float64, y float64) float64 {
	if x == 10 {
		return 10
	} else if x == 20 {
		return 20
	}

	return x + y
}
