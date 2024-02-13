package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	num1 := 15
	num2 := 18
	result := add(num1, num2)
	fmt.Printf("%d +%d = %d\n", num1, num2, result)
}
