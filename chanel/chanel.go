package main

import (
	"fmt"
	"time"
)

func process1(c chan string, data string) {
	c <- data
}

func main() {
	ch := make(chan string, 1)
	go process1(ch, "Data 1")
	fmt.Println(<-ch)
	time.Sleep(5 * time.Second)
}
