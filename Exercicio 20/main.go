package main

import (
	"fmt"
)

func escNome() string { return "HelloWorld" }

func main() {
	nome := escNome()
	fmt.Println(nome)
}
