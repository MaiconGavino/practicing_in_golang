/*
 Um stream vira centenas de go funcs que depois convergem.
        - Dois canais.
        - Uma func manda X números ao primeiro canal.
        - Outra func faz um range deste canal, e para cada ítem lança uma go func que poe o retorno de trabalho() no canal dois.
        - Trabalho() é um timer aleatório pra simular workload.
        - Por fim, range canal dois demonstra os valores.


package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)

	go input(20, canal1)
	go inputOutra(canal1, canal2)

	for v := range canal2 {
		fmt.Println(v)
	}
}
func input(ent int, canal1 chan int) {
	for i := 0; i < ent; i++ {
		canal1 <- i
	}
	close(canal1)
}
func inputOutra(canal1, canal2 chan int) {
	var wg sync.WaitGroup

	for i := range canal1 {
		wg.Add(1)
		go func(x int) {
			canal2 <- trabalho(x)
			wg.Done()
		}(i)
	}
	wg.Wait()
	close(canal2)
}
func trabalho(n int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
	return n
}
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)
	funcoes := 5
	go input(100, canal1)
	go inputOutra(funcoes, canal1, canal2)

	for v := range canal2 {
		fmt.Println(v)
	}
}
func input(ent int, canal1 chan int) {
	for i := 0; i < ent; i++ {
		canal1 <- i
	}
	close(canal1)
}
func inputOutra(funcoes int, canal1, canal2 chan int) {
	var wg sync.WaitGroup
	for i := 0; i < funcoes; i++ {
		wg.Add(1)
		go func() {
			for v := range canal1 {
				canal2 <- trabalho(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(canal2)
}
func trabalho(n int) int {
	time.Sleep(time.Millisecond * 1000) //* time.Duration(rand.Intn(1e3)))
	return n
}
