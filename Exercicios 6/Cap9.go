/*
Exercício 1
- Usando uma literal composta:
     - Crie um array que suporte 5 valores to tipo int
     - Atribua valores aos seus índices
- Utilize range e demonstre os valores do array.
- Utilizando format printing, demonstre o tipo do array.

package main

import "fmt"

func main() {
	example := [5]string{"água", "gelo", "cerveja", "refrigerante", "limão"}
	fmt.Println(example)
	for value, posicao := range example {
		fmt.Println(value, posicao)
	}
	fmt.Printf("%T", example)
}

Exercício 2
- Usando uma literal composta:
    - Crie uma slice de tipo int
    - Atribua 10 valores a ela
- Utilize range para demonstrar todos estes valores.
- E utilize format printing para demonstrar seu tipo.

package main

import "fmt"

func main() {
	slice := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	for key, value := range slice {
		fmt.Println(key, "-", value)
	}
	fmt.Printf("%T", slice)
}

Exercício 3
- Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
    - Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
    - Do quinto ao último item do slice (incluindo o último item!)
    - Do segundo ao sétimo item do slice (incluindo o sétimo item!)
    - Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
    - Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item

package main

import "fmt"

func main() {
	slice := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	fmt.Printf("slice completo %v\n", slice)
	fmt.Printf("Do primeiro ao terceiro item do slice %v\n", slice[:3])
	fmt.Printf("Do quinto ao último item do slice %v\n", slice[4:])
	fmt.Printf("Do segundo ao sétimo item do slice %v\n", slice[1:7])
	fmt.Printf("Do terceiro ao penúltimo item do slice %v\n", slice[2:9])
	fmt.Printf("Obtenha o mesmo resultado acima utilizando a função len() %v\n", slice[2:len(slice)])
}

Exercício 4
- Começando com a seguinte slice:
    - x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
- Anexe a ela o valor 52;
- Anexe a ela os valores 53, 54 e 55 utilizando uma única declaração;
- Demonstre a slice;
- Anexe a ela a seguinte slice:
    - y := []int{56, 57, 58, 59, 60}
- Demonstre a slice x.
package main

import "fmt"

func main() {
	prim := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	prim = append(prim, 52)
	fmt.Println(prim)
	prim = append(prim, 53, 54, 55)
	fmt.Println(prim)
	sec := []int{56, 57, 58, 59, 60}
	prim = append(prim, sec...)
	fmt.Println(prim)

}

Exercício 5
Comece com essa slice:
    - x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
- Utilizando slicing e append, crie uma slice y que contenha os valores:
	- [42, 43, 44, 48, 49, 50, 51]

package main

import "fmt"

func main() {
	first := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	second := append(first[:3], first[6:]...)
	fmt.Println(second)
}

Exercício 6
- Crie uma slice usando make que possa conter todos os estados do Brasil.
	- Os estados: "Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão",
	"Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí",
	"Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina",
	"São Paulo", "Sergipe", "Tocantins"
- Demonstre o len e cap da slice.
- Demonstre todos os valores da slice *sem utilizar range.*

package main

import "fmt"

func main() {
	estados := make([]string, 27, 27)
	estados = []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia",
		"Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul",
		"Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro",
		"Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina",
		"São Paulo", "Sergipe", "Tocantins"}
	for i := 0; i < len(estados); i++ {
		fmt.Println(estados[i])
	}
	fmt.Println("Len = ", len(estados), "\nCap = ", cap(estados))
}

Exercício 7
Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
    - "Nome", "Sobrenome", "Hobby favorito"
- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.

package main

import "fmt"

func main() {
	slice := [][]string{
		[]string{"maicon", "maria", "stella", "elisangela"},
		[]string{"serrão", "eduarda", "mares", "guimarães"},
		[]string{"programar", "enfermegem", "desenhar", "cantar"},
	}
	for key, value := range slice {
		fmt.Println(key, value)
	}
}

Exercício 8
- Crie um map com key tipo string e value tipo []string.
    - Key deve conter nomes no formato sobrenome_nome
    - Value deve conter os hobbies favoritos da pessoa
- Demonstre todos esses valores e seus índices.

package main

import "fmt"

func main() {
	slice := map[string]string{
		"Maicon Serrão":        "desenvolvedor",
		"Elisangela Guimarães": "Business",
		"Stella Mares":         "Desenhista",
	}
	fmt.Printf("Nomes: \t\t\t\t Profissões:\n")
	for key, value := range slice {
		fmt.Println(key, "\t\t\t", value)
	}
}

Exercício 9
Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.
package main

import "fmt"

func main() {
	slice := map[string]string{
		"Maicon Serrão":        "desenvolvedor",
		"Elisangela Guimarães": "Business",
		"Stella Mares":         "Desenhista",
	}
	slice["Rafael Serrão"] = "game"
	fmt.Printf("Nomes: \t\t\t\t Profissões:\n")
	for key, value := range slice {
		fmt.Println(key, "\t\t\t", value)
	}
}

Exercício 10
Utilizando o exercício anterior, remova uma entrada do map e demonstre o map inteiro utilizando range.*/
package main

import "fmt"

func main() {
	slice := map[string]string{
		"Maicon Serrão":        "desenvolvedor",
		"Elisangela Guimarães": "Business",
		"Stella Mares":         "Desenhista",
	}
	slice["Rafael Serrão"] = "game"
	delete(slice, "Maicon Serrão")
	fmt.Printf("Nomes: \t\t\t\t Profissões:\n")
	for key, value := range slice {
		fmt.Println(key, "\t\t\t", value)
	}
}
