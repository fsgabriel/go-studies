package main

import "fmt"

func main() {
	// condicao de emprestimo: mais de 22 anos, empregados, mais de um ano de atividade

	var idade int8 = 23
	var empregado, umAnoAtivo bool = true, true
	var salario int32 = 1000

	if idade > 22 && empregado == true && umAnoAtivo == true {
		if salario < 100000 {
			fmt.Println("Voce tem direito ao emprestimo com juros\n")
		} else {
			fmt.Println("Voce tem direito ao emprestimo sem juros\n")
		}
	}

	if idade < 22 {
		fmt.Println("Nao possui idade maior que 22 anos\n")
	}

	if empregado == false {
		fmt.Println("Voce precisa estar empregado\n")
	}

	if umAnoAtivo == false {
		fmt.Println("Necessita de pelo menos um ano de carteira assinada\n")
	}
}
