package main

import "fmt"

// Uma grande empresa de vendas na web precisa adicionar funcionalidades para adicionar
// produtos aos usuários.Para fazer isso,eles exigem que usuárioseprodutos tenhamo
// mesmo endereço de memória no main do programaenas funções.
// Estruturas necessárias:
//     Usuário:Nome,Sobrenome,E-mail,Produtos(array de produtos).
//     Produto:Nome,preço,quantidade.
// Algumas funções necessárias:
//     Novo produto:recebe nomeepreço,eretorna um produto.
//     Adicionar produto:recebe usuário,produtoequantidade,não retorna nada,adiciona
//     oproduto ao usuário.
//     Deletar produtos:recebe um usuário,apaga os produtos do usuário.

type produto struct {
	nome       string
	preco      float64
	quantidade int16
}

type Usuario struct {
	Nome      string
	SobreNome string
	Email     string
	Produtos  []produto
}

func newProduto(nome string, preco float64, quantidade int16) *produto {
	return &produto{
		nome:       nome,
		preco:      preco,
		quantidade: quantidade,
	}
}

func AddProdutoUsuario(u *Usuario, p *produto) {
	u.Produtos = append(u.Produtos, *p)
}

func DeleteProdutoUsuario(u *Usuario, p string) {
	for i, produto := range u.Produtos {
		if produto.nome == p {
			u.Produtos = append(u.Produtos[:i], u.Produtos[i+1:]...)
		}
	}
}

func main() {
	u := Usuario{
		Nome:      "Gabriel",
		SobreNome: "Silva",
		Email:     "gabriel@gmail.com",
	}

	AddProdutoUsuario(&u, newProduto("Sofa", 1220.0, 1))
	AddProdutoUsuario(&u, newProduto("TV", 2000.0, 1))
	AddProdutoUsuario(&u, newProduto("Copos", 10.0, 10))
	AddProdutoUsuario(&u, newProduto("Teclado", 500.0, 1))
	AddProdutoUsuario(&u, newProduto("Mesa", 800.0, 2))
	AddProdutoUsuario(&u, newProduto("Geladeira", 1500.0, 1))

	DeleteProdutoUsuario(&u, "Copos")

	fmt.Println(u)

	DeleteProdutoUsuario(&u, "Mesa")

	fmt.Println(u)

}
