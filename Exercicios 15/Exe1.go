/*
Praticando goroutines
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	go novasgoroutines(100)
	go goroutine1(1000)
	go goroutine2()
	wg.Wait()
}

func novasgoroutines(input int) {
	for i := 0; i < input; i++ {
		x := i
		go func(input int) {
			fmt.Println("Eu sou a goroutine número:", i+1)
			time.Sleep(20)
			wg.Done()
		}(x)
	}
}

func goroutine1(input int) {
	for i := 0; i < input; i += 3 {
		fmt.Println("Eu sou a goroutine1 número:", i+1)
		time.Sleep(20)
		wg.Done()
	}
}

func goroutine2() {
	for i := 0; i < 500; i++ {
		fmt.Println("goroutine:", i)
		time.Sleep(10)
		wg.Done()
	}
}
