// Exercicio 01
//
// Um escritório de contabilidade precisa acessar os dados de seus funcionários para
// poder realizar diferentes liquidações. Para isso, eles têm todos os detalhes
// necessários em um arquivo .txt.
//
//         1. É necessário desenvolver a funcionalidade para poder ler o arquivo .txt
//         indicado pelo cliente, porém, eles não passaram o arquivo para ser lido pelo
//         nosso programa.
//         2. Desenvolva o código necessário para ler os dados do arquivo chamado
//         “customers.txt” (lembre-se do que você viu sobre o pacote “os”).
//
// Como não temos o arquivo necessário, será obtido um erro e, neste caso, o programa
// terá que exibir um panic ao tentar ler um arquivo que não existe, exibindo a
// mensagem “o arquivo indicado não foi encontrado ou está danificado ”.

// Exercicio 02
//
// O mesmo escritório do exercício anterior, solicita uma funcionalidade para poder
// cadastrar dados de novos clientes. Os dados necessários para cadastrar um cliente são:
//         - Arquivo
//         - Nome e sobrenome
//         - RG
//         - Número de telefone
//         - Endereço

// ● Tarefa 1: O número do arquivo deve ser atribuído ou gerado separadamente e antes
//     da cobrança das despesas restantes. Desenvolva e implemente uma função para
//     gerar um ID que você usará posteriormente para atribuí-lo como um valor a “Arquivo”.
//     Se por algum motivo esta função retornar o valor "nil", deve gerar um panic que
//     interrompa a execução e aborte.

// ● Tarefa 2: Antes de cadastrar um cliente, você deve verificar se o cliente já existe. Para
//     fazer isso, você precisa ler os dados de um arquivo .txt. Em algum lugar do seu
//     código, implemente a função para ler um arquivo chamado “customers.txt” (como no
//     exercício anterior, este arquivo não existe, então qualquer função que tente lê-lo
//     retornará um erro). Você deve lidar adequadamente com esse erro, como vimos até
//     agora. Esse erro deve:

// 1. Gerar um panic;
// 2. Imprimir no console a mensagem: “erro: o arquivo indicado não foi encontrado ou
// está danificado”, e continuar com a execução do programa normalmente.

// Requisitos gerais:
// ● Use recover para recuperar o valor dos panics que podem surgir (exceto na tarefa 1).
// ● Lembre-se de realizar as validações necessárias para cada retorno que possa conter
//     um valor de erro (por exemplo, aqueles que tentam ler arquivos).
//     Crie um erro, personalizando-o ao seu gosto, utilizando qualquer uma das funções
//     que o GO disponibiliza para ele (ele também deve realizar a validação relevante para
//     o caso de erro retornado).

package main

import (
	"math/rand"
	"os"
)

func readFile(path string) *os.File {
	res, err := os.OpenFile(path, os.O_RDONLY, 0777)

	if err != nil {
		panic("o arquivo indicado não foi encontrado ou está danificado ")
	}

	return res
}

func createId() float64 {
	return rand.Float64()
}

func main() {
	readFile("contexto.txt")
}
