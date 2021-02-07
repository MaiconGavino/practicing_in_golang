//Duas go funcs enviando X/2 numeros cada uma pra um canal
//For loop X valores, select case <-x
package main

import (
	"fmt"
)

func main() {
	caseA := make(chan int)
	caseB := make(chan int)
	x := 50
	go func(x int) {
		for i := 0; i < x; i++ {
			caseA <- i
		}
	}(x / 2)
	go func(x int) {
		for i := 0; i < x; i++ {
			caseB <- i
		}
	}(x / 2)

	for i := 0; i < x; i++ {
		select {
		case v := <-caseA:
			fmt.Println("Canal A:", v)
		case v := <-caseB:
			fmt.Println("Canal B:", v)
		}
	}
}
