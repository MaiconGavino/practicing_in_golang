//Chans par, ímpar, quit
// Func send manda números pares pra um canal, ímpares pra outro, e fecha/quit
// Func receive é um select entre os três canais, encerra no quit
package main

import (
	"fmt"
)

func main() {
	par := make(chan int)
	impar := make(chan int)
	quit := make(chan int)
	go num(par, impar, quit)
	apresentar(par, impar, quit)

}
func num(par chan int, impar chan int, quit chan int) {
	for i := 1; i <= 50; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	quit <- 0
}
func apresentar(par chan int, impar chan int, quit chan int) {
	for {
		select {
		case v := <-par:
			fmt.Println("o valor", v, "é par")
		case v := <-impar:
			fmt.Println("o valor", v, "é impar")
		case <-quit:
			return
		}
	}
}
