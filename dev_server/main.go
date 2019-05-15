package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting test server...")
	timeout := 0

	for timeout < 20 {
		fmt.Println("Running")
		time.Sleep(time.Second)
		timeout++
	}

	fmt.Println("Timeout!")
}
