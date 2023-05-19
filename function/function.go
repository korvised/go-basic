package main

import "fmt"

func hello() {
	fmt.Println("Hello World")
}

func sum(a int, b int) {
	result := a + b

	fmt.Println("Result =", result)
}

func getSum(a, b int) int {
	return a + b

}

func main() {
	hello()

	sum(2, 5)

	result := getSum(10, 5)
	fmt.Println("getSum =", result)
}
