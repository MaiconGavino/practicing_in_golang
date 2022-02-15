package main

import "fmt"

func capturaInput() {
	fmt.Printf("Digite seu nome: ")
	var nome string
	fmt.Scan(&nome) //&nome -> [endereço]

	fmt.Printf("Digite sua idade: ")
	var idade int64
	fmt.Scan(&idade) //&idade -> [endereço]

	if idade <= 0 {
		fmt.Printf("Ola %s, você digitou uma idade inválida", nome)
	} else {
		fmt.Printf("Olá %s, eu também tenho %d anos", nome, idade)
	}
}

func main() {
	capturaInput()
}
