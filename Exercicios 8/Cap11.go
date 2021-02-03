/*
### Na prática: exercício #1

- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
    - Nome
    - Sobrenome
    - Sabores favoritos de sorvete
- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
- Solução: https://play.golang.org/p/Pyp7vmTJfY


package main

import "fmt"

type pessoa struct {
	name     string
	surname  string
	iceCream []string
}

func main() {
	pessoa1 := pessoa{
		name:     "maicon",
		surname:  "serrão",
		iceCream: []string{"floresta negra", "flocos", "brigadeiro"}}
	pessoa2 := pessoa{
		name:     "bianca",
		surname:  "kalynic",
		iceCream: []string{"morango", "chocolate", "nutella"}}
	fmt.Println(pessoa1.name, pessoa1.surname)
	for _, v := range pessoa1.iceCream {
		fmt.Println("\t", v)
	}
	fmt.Println(pessoa2.name, pessoa2.surname)
	for _, v := range pessoa2.iceCream {
		fmt.Println("\t", v)
	}
}


### Na prática: exercício #2

- Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
- Demonstre os valores do map utilizando range.
- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
- Solução: https://play.golang.org/p/GLK11Q1_x8y

package main

import "fmt"

type pessoas struct {
	name     map[string]string
	iceCream []string
}

func main() {
	pessoa1 := pessoas{
		name: map[string]string{
			"serrão":  "maicon",
			"gavino":  "rafael",
			"firmino": "erica",
		},
		iceCream: []string{"banana com canele", "nata", "brigadeiro", "morangos"}}
	for key, value := range pessoa1.name {
		fmt.Println(value, key)
		for _, v := range pessoa1.iceCream {
			println("\t", v)
		}
	}
}

### Na prática: exercício #3

- Crie um novo tipo: veículo
    - O tipo subjacente deve ser struct
    - Deve conter os campos: portas, cor
- Crie dois novos tipos: caminhonete e sedan
    - Os tipos subjacentes devem ser struct
    - Ambos devem conter "veículo" como struct embutido
    - O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro"
    - O tipo sedan deve conter um campo bool chamado "modeloLuxo"
- Usando os structs veículo, caminhonete e sedan:
    - Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
    - Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
- Demonstre estes valores.
- Demonstre um único campo de cada um dos dois.
- Solução: https://play.golang.org/p/3eoGb0kxzT


package main

import "fmt"

type vehicle struct {
	door  string
	color string
}
type pickupTruck struct {
	vehicle
	traction bool
}
type sedan struct {
	vehicle
	lux bool
}

func main() {
	caminhonete := pickupTruck{
		vehicle: vehicle{
			door:  "4 portas",
			color: "preta"},
		traction: true,
	}
	fmt.Println(caminhonete)
	CarSedan := sedan{
		vehicle: vehicle{
			door:  "4 portas",
			color: "vermelho",
		},
		lux: false,
	}
	fmt.Println(CarSedan)
	fmt.Println(caminhonete.door)
	fmt.Println(CarSedan.color)
}

### Na prática: exercício #4

- Crie e use um struct anônimo.
- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
- Solução: https://play.golang.org/p/iTGLyH0Ijc & https://play.golang.org/p/h247Kid5adG
*/
package main

import "fmt"

func main() {
	picole := struct {
		tipo  map[string]int
		sabor []string
	}{
		tipo: map[string]int{
			"creme":  1,
			"fruta":  2,
			"itu":    3,
			"paleta": 4,
		}, sabor: []string{"morango", "uva", "melancia", "chocolate branco"}}
	fmt.Println(picole)
}
