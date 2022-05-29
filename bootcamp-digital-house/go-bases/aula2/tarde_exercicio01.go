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

// package main
// import (
// "fmt"
// "time"
// )
// type aluno struct {
// Nome string
// Sobrenome string
// RG string
// DataAdmissao time.Time
// }
// func (a aluno) detalhar() {
// fmt.Println("Nome:", a.Nome)
// fmt.Println("Sobrenome:", a.Sobrenome)
// fmt.Println("RG:", a.RG)
// fmt.Println("DataAdmissao:", a.DataAdmissao)
// }
// func main() {
// a1 := aluno{
// Nome: "João",
// Sobrenome: "da Silva",
// RG: "123456",
// DataAdmissao: time.Date(1980, 01, 15, 0, 0, 0, 0, time.UTC),
// }
// a1.detalhar()
// a2 := aluno{
// Nome: "Maria",
// Sobrenome: "Aparecida",
// RG: "654321",
// DataAdmissao: time.Now().UTC(),
// }
// a2.detalhar()
// }