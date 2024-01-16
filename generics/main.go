package main

import "fmt"

type Number interface {
	int | int64 | float64
}


func sumNumbers[T Number](numbers []T) T {
	var result T
	for i := range numbers {
		result += numbers[i]
	}
	return result
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	result := sumNumbers(numbers)
	fmt.Println("The sum of the number is: ", result)
}