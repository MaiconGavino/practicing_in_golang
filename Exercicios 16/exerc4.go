package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 25
	}()
	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
