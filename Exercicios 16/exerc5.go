package main

import (
	"fmt"
)

func main() {
	canal := make(chan int)

	go input(canal)
	out(canal)

}

func input(canal chan int) {

	for i := 0; i < 100; i++ {
		canal <- i + 1
	}
	close(canal)
}
func out(canal chan int) {
	for v := range canal {
		fmt.Println(v)
	}
}
