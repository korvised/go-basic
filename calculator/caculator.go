package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(promt string) float64 {
	fmt.Printf("%v", promt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message, _ := fmt.Scanf("%v must number only", promt)
		panic(message)
	}

	return value
}

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	return a / b
}

func getOperator() string {
	fmt.Print("operator is (+ - * /):")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func main() {
	value1 := getInput("value1 = ")
	value2 := getInput("value2 = ")

	var result float64

	switch operator := getOperator(); operator {
	case "+":
		result = add(value1, value2)
	case "-":
		result = subtract(value1, value2)
	case "*":
		result = multiply(value1, value2)
	case "/":
		result = divide(value1, value2)
	default:
		panic("wrong operator")
	}

	fmt.Printf("result is %v", result)

}
