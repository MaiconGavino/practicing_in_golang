/*
## 22 – Exercícios: Ninja Nível 10

### Na prática: exercício #1

- Nível 10?! Êita! Parabéns!
- Faça esse código funcionar: https://play.golang.org/p/j-EA6003P0
    - Usando uma função anônima auto-executável

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)
}

 - Usando buffer
- Solução:
*/

package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)
	c <- 42
	fmt.Println(<-c)
}
