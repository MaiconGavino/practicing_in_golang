/*
### Na prática: exercício #1

- Crie um valor e atribua-o a uma variável.
- Demonstre o endereço deste valor na memória.

package main

import "fmt"

func main() {
	example := 3.14159265
	fmt.Println(&example)
}

### Na prática: exercício #2

- Crie um struct "pessoa"
- Crie uma função chamada mudeMe que tenha *pessoa como parâmetro. Essa função deve mudar um valor armazenado no endereço *pessoa.
    - Dica: a maneira "correta" para fazer dereference de um valor em um struct seria (*valor).campo
    - Mas consta uma exceção na documentação. Link: https://golang.org/ref/spec#Selectors
    - "As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method),
    - → x.f is shorthand for (*x).f." ←
    - Ou seja, podemos usar tanto o atalho p1.nome quanto o tradicional (*p1).nome
- Crie um valor do tipo pessoa;
- Use a função mudeMe e demonstre o resultado.
*/

package main

import "fmt"

type person struct {
	name    string
	surname string
	age     int
}

func main() {
	guri := person{
		name:    "Kaio",
		surname: "Fshyking",
		age:     28,
	}
	fmt.Println(guri)
	mudaMe(&guri)
	fmt.Println(guri)

}
func mudaMe(pessoa *person) {
	(*pessoa).name = "tiago"
	pessoa.surname = "Kytting"
	pessoa.age = 13
}
