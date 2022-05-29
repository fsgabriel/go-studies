package main

import (
	"errors"
	"fmt"
)

func main() {
	resultado, err := calcularMediaNotas(10.2, 1.2, 5.7, 8.2, -1)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Media das notas igual: %.1f\n", resultado)
	}
}

func calcularMediaNotas(valores ...float32) (float32, error) {
	var totalNotas float32 = 0.0

	for _, nota := range valores {
		if nota < 0.0 {
			return 0.0, errors.New("nota não pode ser negativa")
		}
		totalNotas += nota
	}

	return totalNotas / float32(len(valores)), nil
}
