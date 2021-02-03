/*
 - Os tipos *profissão1* e *profissão2* contem o tipo *pessoa*
    - Cada um tem seu método *oibomdia()*, e podem dar oi utilizando *pessoa.oibomdia()*
    - Implementam a interface *gente*
    - Ambos podem acessar o função *serhumano()* que chama o método *oibomdia()* de cada *gente*
    - Tambem podemos no método *serhumano()* tomar ações diferentes dependendo do tipo:
        switch pessoa.(type) { case profissão1: fmt.Println(h.(profissão1).valorquesóexisteemprofissão1) [...] }*

package main

import "fmt"

type pessoa struct {
	name    string
	surname string
	age     int
}
type engenheiro struct {
	pessoa
	specialization string
	salary         float64
}
type develop struct {
	pessoa
	linguage string
	area     string
}

func (x engenheiro) nome() {
	fmt.Println("Meu nome é", x.name)
}
func (x develop) nome() {
	fmt.Println("Meu nome é", x.name)
}

type gente interface {
	nome()
}

func serhumano(g gente) {
	g.nome()
	switch g.(type) {
	case engenheiro:
		fmt.Println("e eu tenho especialização em", g.(engenheiro).specialization)
	case develop:
		fmt.Println("e minha linguagens preferidas são", g.(develop).linguage)
	}
}
func main() {
	mreng := engenheiro{
		pessoa: pessoa{
			name:    "Artur",
			surname: "Camargo",
			age:     28,
		},
		specialization: "Controle e Automação",
		salary:         12.396,
	}
	mrdev := develop{
		pessoa: pessoa{
			name:    "Codmago",
			surname: "Gavino",
			age:     24,
		},
		linguage: "C++, Go, C#, Python",
		area:     "Back end",
	}
	//mreng.nome()
	//mrdev.nome()
	serhumano(mreng)
	serhumano(mrdev)
}


- Callback é passar uma função como argumento.
- Exemplo:
    - Criando uma função que toma uma função e um []int, e usa somente os números pares como argumentos para a função.
    - Go Playground:
- Desafio: Crie uma função no programa acima que utilize somente os números *ímpares.*

*/
package main

import "fmt"

func main() {
	t := somenteImpar(soma, []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}...)
	fmt.Println(t)
}
func soma(x ...int) int {
	n := 0
	for _, value := range x {
		n += value
	}
	return n
}
func somenteImpar(f func(x ...int) int, y ...int) int {
	var slice []int
	for _, value := range y {
		if value%2 != 0 {
			slice = append(slice, value)
		}
	}
	total := f(slice...)
	return total
}
