package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main() {

	minhaFunc, _ := operation(minimum)
	averageFunc, _ := operation(average)
	maxFunc, _ := operation(maximum)

	minValue := minhaFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println(minValue, maxValue, averageValue)
}

func operation(operacao string) (func(n ...int) int, error) {
	switch operacao {
	case minimum:
		return minimumFn, nil
	case maximum:
		return maximumFn, nil
	case average:
		return averageFunc, nil
	}

	return func(n ...int) int { return 0 }, errors.New("operacao invalida")
}

func minimumFn(n ...int) int {
	var aux int = n[0]

	for _, value := range n {
		if value < aux {
			aux = value
		}
	}

	return aux
}

func maximumFn(n ...int) int {
	var aux int = n[0]

	for _, value := range n {
		if value > aux {
			aux = value
		}
	}

	return aux
}

func averageFunc(n ...int) int{
	var aux int = 0

	for _, value := range n {
		aux += value
	}

	return aux / len(n)
}
