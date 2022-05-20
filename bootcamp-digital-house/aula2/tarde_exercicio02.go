package main

import "fmt"

type loja struct {
	listProdutos []produtos
}

type produtos struct{
	tipo string
	nome string
	preco float64
	adicionl float64
}

type Produtos interface{
	CalcularCusto()
}

type Ecommerce interface{
	Total()
	Adicionar()
}

func (p *produtos) CalcularCusto() {
	if p.tipo == "Grande"{
		p.adicionl = (p.preco * 0.06) + 2500
	}
	if p.tipo == "Medio"{
		p.adicionl = (p.preco * 0.03)
	}
}

func (l *loja) Total(){
	var total float64

	for _, produto := range l {
		total += l.preco + l.adicionl
	}
}

func NovoProduto(tipo, nome string, preco float64) *produtos{
	return &produtos{
		tipo: tipo, 
		nome: nome, 
		preco: preco,
	}
}

func NovaLoja(listProdutos ...produtos) *loja{
	return &loja{listProdutos: listProdutos}
}

func main(){
	listaProdutos := []produtos{}

	listaProdutos = append(listaProdutos, *NovoProduto("Tec", "Celular", 1000.00))
	listaProdutos = append(listaProdutos, *NovoProduto("Casa", "Sofa", 5000.00))

	loja2 := *NovaLoja(listaProdutos...)

	fmt.Println(loja2)
}