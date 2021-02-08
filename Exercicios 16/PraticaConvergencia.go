/*
 Canais par, ímpar, e converge.
        - Func send manda pares pra um, ímpares pro outro, depois fecha.
        - Func receive cria duas go funcs, cada uma com um for range, enviando dados dos canais par e ímpar pro canal converge. Não esquecer de WGs!
        - Por fim um range retira todas as informações do canal converge.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	par := make(chan int)
	impar := make(chan int)
	converge := make(chan int)

	go input(par, impar)
	go out(par, impar, converge)

	for i := range converge {
		fmt.Println("Valores convergindo:", i)
	}

}

func input(par, impar chan int) {
	for i := 0; i < 50; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	close(par)
	close(impar)
}

func out(par, impar, converge chan int) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for v := range par {
			converge <- v
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for v := range impar {
			converge <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(converge)
}
