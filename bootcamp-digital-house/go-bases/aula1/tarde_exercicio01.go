package main

import "fmt"

func main() {
	palavra := "Brasileiro"

	fmt.Println(len(palavra))

	for i, letra := range palavra {
		fmt.Println(i, " - ", string(letra))
	}
}
