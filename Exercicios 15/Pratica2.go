package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Scanln(&input)
	for i := 0; i < 11; i++ {
		fmt.Printf("| %d x %d = %d \n", i, input, i*input)
	}
}
