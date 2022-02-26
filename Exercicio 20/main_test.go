package main

import (
	"testing"
)

func TestNome(t *testing.T) {
	nome := escNome()
	if nome != "HelloWorld" {
		t.Error("Erro Inesperado", nome)
	}
}
