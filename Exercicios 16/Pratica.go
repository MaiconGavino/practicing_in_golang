package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go myloop(c)
	apresentar(c)
}

func myloop(s chan<- int) {
	for i := 0; i < 11; i++ {
		s <- i
	}
	close(s)
}

func apresentar(a <-chan int) {
	for v := range a {
		fmt.Println("Recevido do canal:", v)
	}
}
