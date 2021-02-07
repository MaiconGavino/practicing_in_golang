//Goroutines and WaitGroups
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	wg.Add(2)
	go loop1()
	go loop2()

	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
}
func loop1() {
	for i := 0; i <= 100; i++ {
		fmt.Println("loop1", i)
		time.Sleep(20)
	}
	wg.Done()
}
func loop2() {
	for i := 0; i <= 100; i++ {
		fmt.Println("loop2", i)
		time.Sleep(20)
	}
	wg.Done()
}
