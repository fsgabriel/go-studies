package main

import (
	"errors"
	"fmt"
)

const (
	dog = "dog"
	cat = "cat"
)

func main() {
	animalDog, err := Animal("dog")
	animalCat, err := Animal("cat")
	animalHamster, err := Animal("hamster")

	if err != nil {
		fmt.Print(err.Error())
	} else {
		var amount float64
		amount += animalDog(5)
		amount += animalCat(8)
		amount += animalHamster(6)

		fmt.Println("Total de racao", amount, "Kg")
	}

}

func Animal(animal string) (func(quantidadeAnimal int) float64, error) {
	switch animal {
	case dog:
		return calcDog, nil
	case cat:
		return calcCat, nil
	}

	return nil, errors.New("animal invalido")
}

func calcDog(quantidadeAnimal int) float64 {
	return float64(quantidadeAnimal) * 10.0
}

func calcCat(quantidadeAnimal int) float64 {
	return float64(quantidadeAnimal) * 5.0
}
