//Func 1 recebe X valores de canal, depois manda qualquer coisa pra chan quit
//Func 2 for infinito, select: case envia pra canal, case recebe de quit
/*package main

import (
	"fmt"
)

func main() {
	canal := make(chan int)
	quit := make(chan int)
	go canalFun(canal, quit)
	apresentar(canal, quit)
}
func canalFun(x chan int, y chan int) {
	for i := 0; i < 10; i++ {
		x <- i
	}
	y <- 0
}

func apresentar(x chan int, y chan int) {
	for {
		select {
		case v := <-x:
			fmt.Println("Saida Canal:", v)
		case <-y:
			return
		}
	}
}
*/
package main

import (
	"fmt"
)

func main() {
	canal := make(chan int)
	quit := make(chan int)
	go canalFun(canal, quit)
	apresentar(canal, quit)
}
func canalFun(canal chan int, quit chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("Saida Canal:", <-canal)
	}
	quit <- 0
}

func apresentar(canal chan int, quit chan int) {
	qualquercoisa := 1
	for {
		select {
		case canal <- qualquercoisa:
			qualquercoisa++
		case <-quit:
			return
		}
	}
}
