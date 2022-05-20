package main

import "fmt"

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main() {

	minhaFunc, _ := operation(minimum)
	// averageFunc, err := operation(average)
	maxFunc, _ := operation(maximum)

	minValue := minhaFunc(2, 3, 3, 4, 10, 2, 4, 5)
	// averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println(minValue, maxValue)
}

func operation(operacao string) (func(n ...int) int, error) {
	switch operacao {
	case minimum:
		return minimumFn, nil
	case maximum:
		return maximumFn, nil
	}

	return func(n ...int) int { return 0 }, nil
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
