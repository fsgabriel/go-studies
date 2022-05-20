package main

import "fmt"

const (
	DESCONTO_SALARIO_ACIMA_50  = 0.17
	DESCONTO_SALARIO_ACIMA_150 = 0.1
)

func main() {
	salario := 10000.00
	// %d %i %f

	fmt.Printf("Desconte de %.2fR$ sobre o salario %.2fR$\n", calcularImposto(float32(salario)), salario)
}

func calcularImposto(s float32) float32 {
	if s >= 150000 {
		return s * DESCONTO_SALARIO_ACIMA_150
	}

	return s * DESCONTO_SALARIO_ACIMA_50
}
