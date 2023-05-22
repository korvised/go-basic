package main

import "fmt"

func add(a, b float64) {
	result := a + b
	fmt.Println("result :", result)
}

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println("i =", i)
	}
}

func deferloop() {
	for j := 0; j < 10; j++ {
		defer fmt.Println("j =", j)
	}
}

func main() {
	// fmt.Println("Welcome to Calculator")
	// defer fmt.Println("End")
	// add(2, 5)
	// defer add(15, 15)
	// defer add(12, 12)
	loop()
	deferloop()
}
