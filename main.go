package main

import "fmt"

func main() {
	result := Sum(1, 1)
	fmt.Println("The result of 1 + 1 is:", result)
}

func Sum(x float64, y float64) float64 {
	return x + y
}
