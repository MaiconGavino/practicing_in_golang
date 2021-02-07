/*
 Esse exercício vai reforçar seus conhecimentos sobre conjuntos de métodos.
    - Crie um tipo para um struct chamado "pessoa"
    - Crie um método "falar" para este tipo que tenha um receiver ponteiro (*pessoa)
    - Crie uma interface, "humanos", que seja implementada por tipos com o método "falar"
    - Crie uma função "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e que chame o método "falar"
    - Demonstre no seu código:
        - Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
        - Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"
- Se precisar de dicas, veja: https://gobyexample.com/interfaces
*/

package main

import (
	"fmt"
)

type Pessoa struct {
	name      string
	linguagem string
}

func (pessoa *Pessoa) fala() {
	fmt.Println("my name is ", pessoa.name, "and i speak", pessoa.linguagem)
}

type humanos interface {
	fala()
}

func dizerAlgumaCoisa(h humanos) {
	h.fala()
}
func main() {
	person := Pessoa{
		name:      "maicon",
		linguagem: "português",
	}
	(&person).fala()
	dizerAlgumaCoisa(&person)
}
