package main

import "fmt"

var umidade, pressao float32
var temperatura int8

func main(){
	umidade, temperatura, pressao = 10.2, 18, 40.3

	fmt.Println(umidade, temperatura, pressao)
}