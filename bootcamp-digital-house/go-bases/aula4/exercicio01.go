// Exercicio 01
// 1. Em sua função “main”, defina uma variável chamada “salario” e atribua um valor do
// tipo “int”.
//
// 2. Crie um erro personalizado com uma struct que implemente “Error()” com a
// mensagem “erro: O salário digitado não está dentro do valor mínimo" em que seja
// disparado quando “salário” for menor que 15.000. Caso contrário, imprima no
// console a mensagem “Necessário pagamento de imposto”.

// Exercicio 02
// Faça a mesma coisa que no exercício anterior, mas reformule o código para que, em vez de
// “Error()”, seja implementado “errors.New()”.

// Exercicio 03
// Repita o processo anterior, mas agora implementando "fmt.Errorf()", para que a mensagem de
// erro receba como parâmetro o valor de "salario", indicando que não atinge o mínimo
// tributável (a mensagem exibida pelo console deve dizer : "erro: o mínimo tributável é 15.000 e
// o salário informado é: [salario]”, onde [salario] é o valor do tipo int passado pelo parâmetro).

// Exercicio 04
// Vamos fazer com que nosso programa seja um pouco mais complexo e útil.
// 1. Desenvolva as funções necessárias para permitir que a empresa calcule:
//     a) Salário mensal de um funcionário segundo a quantidade de horas trabalhadas.
//         - A função receberá as horas trabalhadas no mês e o valor da hora como
//         parâmetro.
//         - Esta função deve retornar mais de um valor (salário calculado e erro).
//         - No caso de o salário mensal ser igual ou superior a R$ 20.000, 10% devem ser
//         descontados como imposto.
//         - Se o número de horas mensais inseridas for menor que 80 ou um número
//         negativo, a função deverá retornar um erro. Deve indicar "erro: o trabalhador
//         não pode ter trabalhado menos de 80 horas por mês".

// 2. Desenvolva o código necessário para cumprir as funcionalidades solicitadas, usando
// “errors.New()”, “fmt.Errorf()” e “errors.Unwrap()”. Não esqueça de realizar as validações dos
// retornos de erro em sua função “main()”.

package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	SALARIO_MINIMO       = 15000
	MINIMO_HORAS_MENSAIS = 80
	SALARIO_DESCONTO     = 20000
	IMPOSTO_DESCONTO     = 0.2
)

type MyError struct {
}

func (e MyError) Error() string {
	return fmt.Sprint("erro: O salário digitado não está dentro do valor mínimo")
}

func Salario(s int64) (string, error) {
	if s < SALARIO_MINIMO {
		return "", errors.New("erro: O salário digitado não está dentro do valor mínimo")
	}

	return "Necessário pagamento de imposto", nil
}

func SalarioFmtError(s int64) (string, error) {
	if s < SALARIO_MINIMO {
		return "", fmt.Errorf("erros: o mínimo tributável é %dR$ e o salário informado é: %dR$\n", SALARIO_MINIMO, s)
	}

	return "Necessário pagamento de imposto", nil
}

func calcSalario(horas int, preco_hora float64) (float64, error) {
	salario := float64(horas) * preco_hora

	if horas < MINIMO_HORAS_MENSAIS {
		return 0, fmt.Errorf("erro: o trabalhador não pode ter trabalhado menos de %d horas por mês\n", MINIMO_HORAS_MENSAIS)
	}

	if salario > SALARIO_DESCONTO {
		salario -= salario * IMPOSTO_DESCONTO
	}

	return salario, nil
}

func main() {
	// sal, err := SalarioFmtError(1000)

	// if err != nil {
	// 	fmt.Printf("%s", err.Error())
	// }

	// fmt.Print(sal)

	salario, err := calcSalario(40, 300)

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	fmt.Printf("Salario do mes: %.2f\n", salario)
}
