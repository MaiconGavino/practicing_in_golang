package main

import (ca
	"fmt"
)

func main() {
	var value1, value2 float64
	fmt.Println("======== Vamos Calcular ===========")
	fmt.Println("Informe o Primeiro Valor: ")
	fmt.Scan(&value1)
	fmt.Println("Informe o Segundo Valor: ")
	fmt.Scan(&value2)
	fmt.Println("======== Escolha a Operação ===========")
	fmt.Println("[0] - Soma")
	fmt.Println("[1] - Substração")
	fmt.Println("[2] - Multiplicação")
	fmt.Println("[3] - Divisão")
	var escolha int
	fmt.Scan(escolha)
	result, err := calc.Div(value1, value2)
	fmt.Println(result)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
