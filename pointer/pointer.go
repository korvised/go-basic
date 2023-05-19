package main

import "fmt"

func zeroValue(i int) {
	i = 0
}

//! define pointer type must use *
func zeroPointer(i *int) {
	*i = 0
}

func main() {
	i := 1
	fmt.Println("i =", i)

	zeroValue(i)
	fmt.Println("i from function zeroValue", i)

	//! to access address of the variable must &
	zeroPointer(&i)
	fmt.Println("i address from function zeroPointer", &i)
	fmt.Println("i value from function zeroPointer", i)
}
