package main

import "fmt"

const (
	SALARIO_CATEGORIA_A = 3000.0
	SALARIO_CATEGORIA_B = 1500.0
	SALARIO_CATEGORIA_C = 1000.0

	BONUS_HORA_CATEGORIA_A = 0.5
	BONUS_HORA_CATEGORIA_B = 0.2

	HORAS_EXTRAS = 160
)

func main() {
	fmt.Printf("Salario total: %.2f\n", calcularSalario("B", 159))
}

func calcularSalario(categoria string, horas float64) float64 {
	var salarioBruto float64

	if categoria == "A" {
		salarioBruto = SALARIO_CATEGORIA_A * horas

		if horas > HORAS_EXTRAS {
			salarioBruto += salarioBruto * BONUS_HORA_CATEGORIA_A
		}
	}

	if categoria == "B" {
		salarioBruto = SALARIO_CATEGORIA_B * horas

		if horas > HORAS_EXTRAS {
			salarioBruto += salarioBruto * BONUS_HORA_CATEGORIA_B
		}
	}

	if categoria == "C" {
		salarioBruto = SALARIO_CATEGORIA_A * horas
	}

	return salarioBruto
}
