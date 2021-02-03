/*
 Exercício:
    - Crie uma função que retorne um int
    - Crie outra função que retorne um int e uma string
    - Chame as duas funções
    - Demonstre seus resultados

package main

import "fmt"

func main() {
	var fatorial int
	fmt.Scanln(&fatorial)
	x := fat(fatorial)
	fmt.Println(resp(x))

}
func fat(fator int) int {
	result := 1
	if fator <= 0 {
		result = 0
		return result
	} else {
		for i := 1; i <= fator; i++ {
			result = result * i
		}
		return result
	}
}
func resp(fator int) (string, int) {
	var name string
	if fator == 0 {
		name = "O fatorial de 0! é:"
		return name, fator
	} else {
		name = "Esse fatorial é:"
		return name, fator
	}
}

### Na prática: exercício #2

- Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
- Passe um valor do tipo slice of int como argumento para a função;
- Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
- Passe um valor do tipo slice of int como argumento para a função.

package main

import "fmt"

func main() {
	x := []int{1, 5, 7, 6, 5, 3, 8, 9}
	y := []int{42, 24, 63, 74, 23, 56}
	fmt.Println(variatico(x...))
	fmt.Println(argum(y))
}
func variatico(x ...int) int {
	sum := 0
	for _, value := range x {
		sum += value
	}
	return sum
}
func argum(x []int) int {
	sum := 0
	for _, value := range x {
		sum += value
	}
	return sum
}

### Na prática: exercício #3

- Utilize a declaração defer de maneira que demonstre que sua execução só ocorre ao final do contexto ao qual ela pertence.

package main

import "fmt"

func main() {
	defer fmt.Println("Com a função Defer")
	fmt.Println("Sem a função Defer")
	fmt.Println("Deveria ser a ultima a ser apresentada")
}

### Na prática: exercício #4

- Crie um tipo struct "pessoa" que contenha os campos:
    - nome
    - sobrenome
    - idade
- Crie um método para "pessoa" que demonstre o nome completo e a idade;
- Crie um valor de tipo "pessoa";
- Utilize o método criado para demonstrar esse valor.

package main

import "fmt"

type pessoa struct {
	name    string
	surname string
	age     int
}

func (escrever pessoa) escrever() {
	fmt.Println("Meu nome é ", escrever.name, "e meu sobrenome é ", escrever.surname, "eu declaro que tenho", escrever.age, "anos de idade")
}
func main() {
	pessoa1 := pessoa{
		name:    "maicon",
		surname: "serrão",
		age:     23,
	}
	pessoa1.escrever()
}


### Na prática: exercício #5

- Crie um tipo "quadrado"
- Crie um tipo "círculo"
- Crie um método "área" para cada tipo que calcule e retorne a área da figura
    - Área do círculo: 2 * π * raio
    - Área do quadrado: lado * lado
- Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
- Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
- Crie um valor de tipo "quadrado"
- Crie um valor de tipo "círculo"
- Use a função "info" para demonstrar a área do "quadrado"
- Use a função "info" para demonstrar a área do "círculo"

package main

import "fmt"

type circle struct {
	raio float64
}
type square struct {
	lado float64
}

func (areaCircle circle) area() {
	value := (areaCircle.raio * areaCircle.raio) * 3.1415
	fmt.Println("A área do circulo é:", value)
}
func (areaSquare square) area() {
	value := areaSquare.lado * areaSquare.lado
	fmt.Println("A área do quadrado é:", value)

}

type info interface {
	area()
}

func medida(i info) {
	i.area()
}

func main() {
	var raioEnt, ladoEnt float64
	fmt.Scanln(&raioEnt)
	fmt.Scanln(&ladoEnt)
	raio := circle{
		raio: raioEnt,
	}
	lado := square{
		lado: ladoEnt,
	}
	medida(raio)
	medida(lado)
}
*/
/*
### Na prática: exercício #6

- Crie e utilize uma função anônima.
package main

import (
	"fmt"
	"math"
)

func main() {
	func(raio float64) {
		fmt.Println(raio * raio * math.Pi)
	}(7.0)
}

### Na prática: exercício #7

- Atribua uma função a uma variável.
- Chame essa função.

package main

import (
	"fmt"
	"math"
)

func main() {
	x := func(raio float64) {
		fmt.Println(math.Pow(raio, 2) * math.Pi)
	}
	x(7.0)
}

### Na prática: exercício #8

- Crie uma função que retorna uma função.
- Atribua a função retornada a uma variável.
- Chame a função retornada.

package main

import "fmt"

func main() {
	saida := sum()(12, 18)
	fmt.Println(saida)
}
func sum() func(x, y int) int {
	soma := 0
	return func(x, y int) int {
		soma += x * y
		return soma
	}
}

### Na prática: exercício #9

- Callback: passe uma função como argumento a outra função.

package main

import (
	"fmt"
)

func main() {
	saida(entrada())
}
func entrada() int {
	var test int
	fmt.Scanln(&test)
	return test
}
func saida(x int) {
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", i, x, i*x)
	}
}

### Na prática: exercício #10

- Demonstre o funcionamento de um closure.
- Ou seja: crie uma função que retorna outra função, onde esta outra função faz uso de uma variável alem de seu scope.

package main

import "fmt"

var soma int

func main() {
	said := entrada()
	fmt.Println(said())
}
func entrada() func() int {
	var ent int
	fmt.Scanln(&ent)
	soma = 0
	return func() int {
		for i := 0; i < ent; i++ {
			if i%2 == 0 {
				soma += i
			}
		}
		return soma
	}
}


exemplo*/

package main

import (
	"fmt"
	"math"
)

type cone struct {
	lightning float64
	heigth    float64
}
type cube struct {
	edge float64
}
type cylinder struct {
	lightning float64
	heigth    float64
}

func (value cone) volume() {
	vol := (math.Pi * math.Pow(value.lightning, 2) * value.heigth) / 3
	fmt.Printf("O volume do cone de raio %f e altura %f é: %f\n", value.lightning, value.heigth, vol)
}
func (value cube) volume() {
	vol := math.Pow(value.edge, 3)
	fmt.Printf("O volume do cubo de aresta %f é: %f\n", value.edge, vol)
}
func (value cylinder) volume() {
	vol := math.Pi * value.lightning * value.heigth
	fmt.Printf("O volume do cilindro de raio %f e altura %f é: %f\n", value.lightning, value.heigth, vol)
}

type vol interface {
	volume()
}

func input(i vol) {
	i.volume()
}
func main() {
	conee := cone{
		lightning: 7.65,
		heigth:    6.35,
	}
	cubo := cube{
		edge: 4.78,
	}
	cilindro := cylinder{
		lightning: 0.75,
		heigth:    5.69,
	}
	input(conee)
	input(cubo)
	input(cilindro)
}
