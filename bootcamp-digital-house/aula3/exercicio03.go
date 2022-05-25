package main

// Uma empresa de mídia social precisa implementar uma estrutura de usuários com funções
// que acrescentem informaçõesàestrutura.Para otimizareeconomizar memória,eles exigem
// queaestrutura de usuários ocupeomesmo lugar na memória paraoprograma principale
// para as funções:
//     Aestrutura deve possuir os seguintes campos:Nome,Sobrenome,idade,emaile
//      senha
// Edevem implementar as seguintes funcionalidades:
//      mudaronome:me permite mudaronomeesobrenome
//      mudaraidade:me permite mudaraidade
//      mudare-mail:me permite mudaroe-mail
//      mudarasenha:me permite mudarasenha

import (
	"fmt"
)

type Person struct{
	Nome string
	Sobrenome string
	Idade int
	Email string
	Senha string
}

func (p *Person) ChangeName(s string){
	p.Nome = s
}

func (p *Person) ChangeSobrenome(s string){
	p.Sobrenome = s 
}

func (p *Person) ChangeIdade(i int){
	p.Idade = i
}

func (p *Person) ChangeSenha(s string){
	p.Senha = s
}

func (p *Person) ChangeEmail(s string){
	p.Email = s
}

func main(){
	p := Person{
		Nome: "Gabriel",
		Sobrenome: "Silva",
		Idade: 22,
		Email: "gabriel@gmail.com",
		Senha: "22222",
	}

	fmt.Println(p)

	p.ChangeIdade(30)
	p.ChangeName("Maria")
	p.ChangeSobrenome("Ferreira")
	p.ChangeSenha("1111")
	p.ChangeEmail("maria@gmail.com")

	fmt.Println(p)
}