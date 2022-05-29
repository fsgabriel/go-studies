// Dependendo do tema escolhido, gere um JSON que atenda as seguintes chaves de acordo
// com o tema.
// Os produtos variam por id, nome, cor, preço, estoque, código (alfanumérico), publicação
// (sim-não), data de criação.
// Os usuários variam por id, nome, sobrenome, e-mail, idade, altura, ativo (sim-não), data de
// criação.
// Transações: id, código da transação (alfanumérico), moeda, valor, emissor (string), receptor
// (string), data da transação.
//         1. Dentro da pasta go-web crie um arquivo theme.json, o nome tem que ser o tema
//         escolhido, ex: products.json.

//         2. Dentro dele escrevi um JSON que permite ter uma matriz de produtos, usuários ou
//         transações com todas as suas variantes.

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var id int64 = 1

type Products struct {
	Id      int64   `json:"id"`
	Nome    string  `json:"nome"`
	Cor     string  `json:"cor"`
	Preco   float64 `json:"preco"`
	Estoque bool    `json:"estoque"`
	Codigo  string  `json:"codigo"`
}

func createID() int64 {
	id += 1
	return id
}

func main2() {
	p := []Products{
		{
			Id:      createID(),
			Nome:    "TV",
			Cor:     "Prata",
			Preco:   1500,
			Estoque: true,
			Codigo:  "dasdas89d78a9s",
		},
		{
			Id:      createID(),
			Nome:    "Sofa",
			Cor:     "Marrom",
			Preco:   1200,
			Estoque: true,
			Codigo:  "pidsf98s09kj",
		},
		{
			Id:      createID(),
			Nome:    "Teclado",
			Cor:     "Prata",
			Preco:   200,
			Estoque: false,
			Codigo:  "odjoaksjda098",
		},
		{
			Id:      createID(),
			Nome:    "Mouse",
			Cor:     "Branco",
			Preco:   20,
			Estoque: false,
			Codigo:  "mlkamds0i9",
		},
	}

	data, err := json.Marshal(p)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	os.WriteFile("./products.json", data, 0777)
}
