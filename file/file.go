package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("account.xlsx")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	count := 1
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println("line", line)
		fmt.Printf("Line %d : %s \n", count, line)

		count++
	}

}
