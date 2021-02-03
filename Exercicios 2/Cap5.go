/*
#EXERCICIO 1
- Escreva um programa que mostre um número em decimal, binário, e hexadecimal.


package main
import	"fmt"
func main(){
	var entrada int
	fmt.Scanf("%d", &entrada)
	fmt.Println("Decimal \t \t Binário \t \t Hexadecimal")
	fmt.Printf("%d\t \t \t %b \t \t \t %#x", entrada, entrada, entrada)
}

#EXERCICIO 2
 Escreva expressões utilizando os operadores, e atribua seus valores a variáveis.
 = != <= >= > <
- Demonstre estes valores.


package main
import "fmt"
func main()  {
	var entA, entB int
	fmt.Scanln( &entA)
	fmt.Scanln( &entB)
	if entA == entB {
		fmt.Printf("%d é igual a %d\n", entA, entB)
	}
	if entA != entB {
		fmt.Printf("%d é diferente de %d\n", entA, entB)
	}
	if entA <= entB{
		fmt.Printf("%d é menor ou igual a %d\n", entA, entB)
	}
	if entA >= entB{
		fmt.Printf("%d é maior ou igual a %d\n", entA, entB)
	}
	if entA > entB{
		fmt.Printf("%d é maior que %d\n", entA, entB)
	}
	if entA < entB{
		fmt.Printf("%d é menor que %d\n", entA, entB)
	}
}

#EXERCICIO 3
- Crie constantes tipadas e não-tipadas.
- Demonstre seus valores.

package main
import "fmt"

const consTip int = 25
const consNtip = 25

func main()  {
		fmt.Printf(consTip)
		fmt.Println(consNtip)
}

#EXERCICIO 4
 Crie um programa que:
    - Atribua um valor int a uma variável
    - Demonstre este valor em decimal, binário e hexadecimal
    - Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
	- Demonstre esta outra variável em decimal, binário e hexadecimal


package main
import "fmt"
func main()  {
	var entrada, desloc int
	fmt.Scanln(&entrada)
	fmt.Println("Sem Descolamento")
	fmt.Println("Decimal \t\t Binário \t\t Hexadecimal")
	fmt.Printf("%d \t\t\t %b \t\t\t %#x\n", entrada, entrada, entrada)
	desloc = entrada << 1
	fmt.Println("Com Descolamento")
	fmt.Println("Decimal \t\t Binário \t\t Hexadecimal")
	fmt.Printf("%d \t\t\t %b \t\t\t %#x", desloc, desloc, desloc)
}

#EXERCICIO 5
- Crie uma variável de tipo string utilizando uma raw string literal.
- Demonstre-a.


package main
import "fmt"
func main()  {
	example := `String
				 of
				  type
				   raw
				     string
					   literal`
	fmt.Println(example)
}

#EXERCICIO 6
- Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
- Demonstre estes valores.
*/
package main

import "fmt"

const (
	prim = iota + 2022
	seg  = iota + 2022
	ter  = iota + 2022
	quar = iota + 2022
)

func main() {
	fmt.Printf("%v, %v, %v, %v", prim, seg, ter, quar)
}
