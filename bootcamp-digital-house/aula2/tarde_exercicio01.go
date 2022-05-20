package main

import "fmt"

type Aluno struct{
	Nome string
	Sobrenome string
	RG int64
	Data string
}

func (aluno *Aluno) Detalhamento(){
	fmt.Printf("Nome completo: %s %s, RG: %.7d, Data de nascimento: %s\n", aluno.Nome, aluno.Sobrenome, aluno.RG, aluno.Data)
}

func main(){
	p := Aluno{
		Nome: "Gabriel", 
		Sobrenome: "Silva", 
		RG: 00000000000, 
		Data: "01/10/1999",
	}

	p.Detalhamento()
}