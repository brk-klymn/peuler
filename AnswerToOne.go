package main

import "fmt"

func main() {
	var sum int

	sum = 0

	for i := 0; i < 1000; i++ {
		if i%5 == 0 {
			sum = sum + i
		} else if i%3 == 0 {
			sum = sum + i
		}
	}

	fmt.Println("The answer is ", sum)
}
