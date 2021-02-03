/*
Exercicío 1
Põe na tela: todos os números de 1 a 10000.

package main

import "fmt"

func main() {
	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}
}

Exercicío 2
Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.
- Por exemplo:
    65
        U+0041 'A'
        U+0041 'A'
        U+0041 'A'
    66
        U+0042 'B'
        U+0042 'B'
        U+0042 'B'
	...e por aí vai.

package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%#U\t\t", i)
		}
		fmt.Println()
	}
}

Exercicío 3
Crie um loop utilizando a sintaxe: for condition {}
- Utilize-o para demonstrar os anos desde que você nasceu.
package main

import "fmt"

func main() {
	nascimento := 1997
	anoAtual := 2021
	for nascimento <= anoAtual {
		fmt.Printf("Ano de %v \n", nascimento)
		nascimento++
	}
}

Exercicío 4
Crie um loop utilizando a sintaxe: for {}
- Utilize-o para demonstrar os anos desde que você nasceu.

func main() {

	for i := 1997; i <= 2021; i++ {
		fmt.Printf("Ano de %v \n", i)
	}
}

Exercicío 5
Demonstre o resto da divisão por 4 de todos os números entre 10 e 100

package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("O resto de %v dividido por 4 é  %v\n", i, i%4)
	}
}

Exercicío 7
Crie um programa que demonstre o funcionamento da declaração if.
Utilizando a solução anterior, adicione as opções else if e else.
package main

import "fmt"

func main() {
	var entrada int
	fmt.Scanf("%d", &entrada)
	if entrada%2 == 0 {
		fmt.Printf("O número %v é par", entrada)
	} else {
		fmt.Printf("O número %v é impar", entrada)
	}
}

Exercicío 8
- Crie um programa que utilize a declaração switch, sem switch statement especificado.

package main

import "fmt"

func main() {
	var inputName string
	fmt.Scanf("%s", &inputName)
	switch {
	case inputName == "joão":
		fmt.Printf("O clientre é o João")
	case inputName == "gabriel":
		fmt.Printf("O cliente é o Gabriel")
	case inputName == "ana":
		fmt.Printf("A cliente é a Ana")
	case inputName == "maicon":
		fmt.Printf("O cliente é o Maicon")
	}
}

Exercicío 9
- Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".
*/
package main

import "fmt"

func main() {
	var esporteFavorito string
	fmt.Scanf("%s", &esporteFavorito)
	switch esporteFavorito {
	case "futebol":
		fmt.Printf("O seu esporte favorito é futebol")
	case "volei":
		fmt.Printf("O seu esporte favorito é volei")
	case "basquete":
		fmt.Printf("O seu esporte favorito é basquete")
	case "handebol":
		fmt.Printf("O seu esporte favorito é handebol")
	case "atletismo":
		fmt.Printf("O seu esporte favorito é atletismo")
	}
}
