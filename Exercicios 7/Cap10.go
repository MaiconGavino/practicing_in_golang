/*
Trabalhando com STRUCTs
package main

import "fmt"

type cliente struct {
	name        string
	surname     string
	reservation bool
}

func main() {
	var nome, sobrenome string
	var reservas bool
	i := 0
	for true {
		fmt.Scanln(&nome)
		fmt.Scanln(&sobrenome)
		fmt.Scanln(&reservas)
		cliente1 := cliente{
			name:        nome,
			surname:     sobrenome,
			reservation: reservas,
		}
		fmt.Println(cliente1)
		if i == 3 {
			break
		}
		i++
	}
}


Trabalhando com STRUCTs Embutidos
package main

import "fmt"

type pess struct {
	name    string
	surname string
}
type profissao struct {
	profession string
	salary     int
}

func main() {
	var nome, sobrenome, prof string
	var salario int
	var test bool
	for true {
		fmt.Scanln(&nome)
		fmt.Scanln(&sobrenome)
		fmt.Println("Possui profissão? true or false")
		fmt.Scanln(&test)
		if test {
			fmt.Scanln(&prof)
			fmt.Scanln(&salario)
		}
		if test {
			pessoa2 := profissao{
				profession: prof,
				salary:     salario,
			}
			fmt.Println(pessoa2)
		}
		//fmt.Println(pessoa1)
		pessoa1 := pess{
			name:    nome,
			surname: sobrenome,
		}
		fmt.Println(pessoa1)
		if nome == "fim" {
			break
		}
	}
}
*/
package main

import "fmt"

type pessoa struct {
	name    string
	surname string
}
type profissao struct {
	pessoa
	profession string
	salary     int
}

func main() {
	var nome, sobrenome, prof string
	var salario int
	var test bool
	for true {
		fmt.Scanln(&nome)
		fmt.Scanln(&sobrenome)
		fmt.Println("Possui profissão? true or false")
		fmt.Scanln(&test)
		if test {
			fmt.Scanln(&prof)
			fmt.Scanln(&salario)
		}
		if test {
			caso := profissao{
				pessoa: pessoa{
					name:    nome,
					surname: sobrenome,
				},
				profession: prof,
				salary:     salario,
			}
			fmt.Println(caso)
		} else {
			caso := pessoa{
				name:    nome,
				surname: sobrenome,
			}
			fmt.Println(caso)
		}
	}
}
