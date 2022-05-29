package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Exercício 2 Ler arquivo
// 		Amesma empresa precisa leroarquivo armazenado,para isso exige que:
// Seja impresso na tela os valores tabelados,com título(àesquerda paraoIDeà
// direita paraoPreçoeQuantidade),opreço,aquantidadeeabaixo do preçoototal
// deve ser exibido(somando preço por quantidade)

// Exemplo:
// ID				Preco				Quantidade
// 111223			30012.00			1
// 444321			1000000.00			4
// 434321			50.50				1

func printfCSV(path string){
	byteMsg, err := os.ReadFile(path)

	if err != nil {
		log.Printf(err.Error())
	}

	str := string(byteMsg)

	s := strings.Split(str, ";")

	fmt.Println(s)
}

type Info struct{
	Id int8
	Produto string
	Preco float32
	Quantidade int32
}

func (i Info) detalhamento() string{
	return fmt.Sprintf("%d, %s, %.2f, %d;", i.Id, i.Produto, i.Preco, i.Quantidade)
}

func (i Info) SaveCsv(path string){
	os.WriteFile(path, []byte(i.detalhamento()), 0644)
}