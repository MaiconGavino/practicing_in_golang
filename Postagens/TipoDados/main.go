// package main

// import (
// 	"fmt"
// )

// var (
// 	nome        string  = "maicon"
// 	ligado      bool    = true
// 	ano         int32   = 2022
// 	temperatura float32 = 24.7
// )

// func main() {
// 	fmt.Printf("valor da variável nome: %s\n", nome)
// 	fmt.Printf("valor da variável ligado: %t\n", ligado)
// 	fmt.Printf("valor da variável ano: %d\n", ano)
// 	fmt.Printf("valor da variável temperatura: %.1f\n", temperatura)
// }

// package main

// import "fmt"

// type Person struct {
// 	nome  string
// 	idade int32
// }

// func main() {
// 	pessoa1 := Person{"Maicon Gavino", 24}
// 	fmt.Printf("Nome: %s\nIdade: %d\n\n", pessoa1.nome, pessoa1.idade)
// }

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v - ", i)
	}
}
