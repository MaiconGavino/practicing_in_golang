package main

import (
	"https://github.com/MaiconGavino/practicing_in_golang/tree/master/Exercicios%2019/calc/calc"
	"fmt"
)

func main() {
	valor1, valor2 := imput()
	esc := escolhas()
	operacao(valor1, valor2, esc)
}
func imput() (float64, float64) {
	var value1, value2 float64
	fmt.Println("======== Vamos Calcular ===========")
	fmt.Println("Informe o Primeiro Valor: ")
	fmt.Scan(&value1)
	fmt.Println("Informe o Segundo Valor: ")
	fmt.Scan(&value2)
	return value1, value2
}
func escolhas() int {
	fmt.Println("======== Escolha a Operação ===========")
	fmt.Println("[0] - Soma")
	fmt.Println("[1] - Subtração")
	fmt.Println("[2] - Multiplicação")
	fmt.Println("[3] - Divisão")
	var escolha int
	fmt.Scan(escolha)
	return escolha
}

func operacao(valor1, valor2 float64, escolha int) {

	switch {
	case escolha == '0':
		result, _ := calc.Soma(value1, value2)
		fmt.Println("A soma de %f.2 e %f.2 é: ", valor1, valor2, result)

	case escolha == '1':
		result, _ := calc.Subtracao(value1, value2)
		fmt.Println("A subtração de %f.2 e %f.2 é: ", valor1, valor2, result)

	case escolha == '2':
		result, _ := calc.Multiplicacao(value1, value2)
		fmt.Println("A multiplicação de %f.2 e %f.2 é: ", valor1, valor2, result)

	case escolha == '3':
		result, err := calc.Divisao(value1, value2)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("A divisão de %f.2 e %f.2 é: ", valor1, valor2, result)
		}

	default:
		fmt.Println("Escolha Invalida")
	}
}
