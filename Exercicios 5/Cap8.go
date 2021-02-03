/*
Exercício: tente acessar todos os itens de uma slice *sem* utilizar range.

package main

import "fmt"

func main() {
	example := []string{"Pizza", "Hamburgue", "Batata-frita", "Refrigerante", "Porção"}
	for i := 0; i < len(example); i++ {
		fmt.Printf("%v - %v\n", i, example[i])
	}
}

Exercício : Slice: anexando a uma slice

package main

import "fmt"

func main() {
	slice1 := []int{1, 3, 5, 7, 9, 11}
	slice2 := []int{2, 4, 6, 8, 10, 12}
	fmt.Println(slice1)
	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)
}

Exercício : Deletando de uma slice

package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice1 = append(slice1[:3], slice1[6:]...)
	fmt.Println(slice1)
}

Exercício : Manipule um slice com make
package main

import "fmt"

func main() {
	slice := make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		if i >= len(slice) {
			slice = append(slice, i*2)
		}
		slice[i] = i * 2
	}
	fmt.Println(slice, len(slice), cap(slice))
	slice = append(slice, 22)
	fmt.Println(slice, len(slice), cap(slice))
}

Exercício : Manipule um slice multi-dimensionais
package main

import "fmt"

func main() {
	example := [][]int{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
	}
	fmt.Println(example[1][5])
}

Exercício : Manipule a função MAP
package main

import "fmt"

func main() {
	agenda := map[string]int{
		"maicon": 992181524,
		"rafael": 993363652,
		"arthur": 965862351,
	}
	fmt.Println(agenda)
	//Adicionando um elemento ao map
	agenda["carlos"] = 993326524
	fmt.Println(agenda)
	fmt.Println(agenda["maicon"])
	//comma ok idiom quando não tem o valor no map
	if input, ok := agenda["exemplo"]; !ok {
		fmt.Println("Nome não encontrado!")
	} else {
		fmt.Println(input)
	}
}


Exercício : Manipule um range in Map
*/
package main

import "fmt"

func main() {
	example := map[string]int{
		"maicon": 992181524,
		"rafael": 993363652,
		"arthur": 965862351,
	}
	fmt.Println(example)
	for key, value := range example {
		fmt.Println(key, value)
	}
	//deletando elemento do map
	delete(example, "arthur")
	fmt.Println(example)
}
