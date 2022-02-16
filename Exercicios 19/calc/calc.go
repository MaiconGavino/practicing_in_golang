package calc

import "errors"

func Divisao(value1, value2 float64) (float64, error) {
	if value2 == 0 {
		return 0, errors.New("Não existe divisão por zero")
	}
	return value1 / value2, nil
}

func Multiplicacao(value1, value2 float64) (float64){return value1 * value2}
func Soma(value1, value2 float64) (float64){return value1 + value2}
func Subtracao(value1, value2 float64) (float64){return value1 - value2}