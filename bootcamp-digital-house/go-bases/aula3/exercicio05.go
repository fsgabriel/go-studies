package main

import (
	"fmt"
	"sync"
)

// Uma empresa nacional é responsável pela venda de produtos, serviços e manutenção.
// Para isso, eles precisam realizar um programa que seja responsável por calcular o preço total
// dos Produtos, Serviços e Manutenção. Devido à forte demanda e para otimizar a velocidade,
// eles exigem que o cálculo da soma seja realizado em paralelo por meio de 3 go routines.

// Precisamos de 3 estruturas:
// - Produtos: nome, preço, quantidade.
// - Serviços: nome, preço, minutos trabalhados.
// - Manutenção: nome, preço.
// Precisamos de 3 funções:
// - Somar Produtos: recebe um array de produto e devolve o preço total (preço *
// quantidade).
// - Somar Serviços: recebe uma array de serviço e devolve o preço total (preço * média
// hora trabalhada, se ele não vier trabalhar por 30 minutos, ele será cobrado como se
// tivesse trabalhado meia hora).
// - Somar Manutenção: recebe um array de manutenção e devolve o preço total.

// Os 3 devem ser executados concomitantemente e ao final o valor final deve ser mostrado na
// tela (somando o total dos 3).

const (
	TEMPO_MINIMO float64 = 30
)

type Produto struct {
	Nome       string
	Preco      float64
	Quantidade int16
}

type Servicos struct {
	Nome               string
	Preco              float64
	MinutosTrabalhados int16
}

type Manutencao struct {
	Nome  string
	Preco float64
}

func PrecoTotalProduto(p []Produto, ch chan string, wg *sync.WaitGroup) float64 {
	defer wg.Done()
	var total float64
	for _, produto := range p {
		total += (produto.Preco * float64(produto.Quantidade))
	}

	ch <- fmt.Sprintf("Soma total preco dos produtos: %.2fR$\n", total)

	return total
}

func PrecoTotalServico(a []Servicos, ch chan string, wg *sync.WaitGroup) float64 {
	defer wg.Done()
	var total float64
	for _, servico := range a {
		if servico.MinutosTrabalhados <= 30 {
			total += (servico.Preco * TEMPO_MINIMO)

		} else {
			total += (servico.Preco * float64(servico.MinutosTrabalhados))
		}
	}

	ch <- fmt.Sprintf("Preco total de servicos por tempo de trabalho: %.2fR$\n", total)

	return total
}

func PrecoTotalManutencao(m []Manutencao, ch chan string, wg *sync.WaitGroup) float64 {
	defer wg.Done()
	var total float64
	for _, manutencao := range m {
		total += manutencao.Preco
	}

	ch <- fmt.Sprintf("Preco total das manutencoes: %.2f\n", total)

	return total
}

func main() {
	ch := make(chan string, 3)

	var wg sync.WaitGroup

	wg.Add(3)

	arrProdutos := []Produto{
		{
			Nome:       "Copos",
			Preco:      10.0,
			Quantidade: 100,
		},
		{
			Nome:       "Sofa",
			Preco:      2000.0,
			Quantidade: 2,
		}, {
			Nome:       "Monitor",
			Preco:      1500.0,
			Quantidade: 2,
		}, {
			Nome:       "Monitor",
			Preco:      1500.0,
			Quantidade: 2,
		},
	}

	arrServicos := []Servicos{
		{
			Nome:               "Envio",
			Preco:              20.0,
			MinutosTrabalhados: 20,
		}, {
			Nome:               "Envio",
			Preco:              50.0,
			MinutosTrabalhados: 50.0,
		}, {
			Nome:               "Envio",
			Preco:              70.0,
			MinutosTrabalhados: 20,
		},
	}

	arrManutencao := []Manutencao{
		{
			Nome:  "Computador",
			Preco: 5.0,
		},
		{
			Nome:  "Rede",
			Preco: 100.0,
		},
	}

	go PrecoTotalManutencao(arrManutencao, ch, &wg)
	go PrecoTotalProduto(arrProdutos, ch, &wg)
	go PrecoTotalServico(arrServicos, ch, &wg)

	wg.Wait()

	close(ch)

	for el := range ch {
		fmt.Println(el)
	}

}
