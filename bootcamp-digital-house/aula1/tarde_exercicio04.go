package main

import (
	"debug/plan9obj"
	"fmt"
)

func main() {
	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	var totalMaiorVintoUm int = 0

	employees["Frederico"] = 25

	delete(employees, "Pedro")

	for _, idade := range employees {
		if idade > 21 {
			totalMaiorVintoUm += 1
		}
	}

	fmt.Println("Total de funcionarios com mais de 21 anos: ", totalMaiorVintoUm, "\n")

	fmt.Println(employees, "\n")
}
