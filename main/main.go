package main

import (
	"../p904"
	"fmt"
)

const cur = "904t"

func main() {

	switch cur {
	case "904":
		test_p904()
	case "904t":
		test_p904t()

	}
}

func test_p904() {
	var input []int

	input = []int{1, 2, 1}
	fmt.Println("Input : ", input)
	fmt.Println("Answer: ", 3)
	fmt.Println("Output: ", p904.TotalFruit(input), "\n")

	input = []int{0, 1, 2, 2}
	fmt.Println("Input : ", input)
	fmt.Println("Answer: ", 3)
	fmt.Println("Output: ", p904.TotalFruit(input), "\n")

	input = []int{1, 2, 3, 2, 2}
	fmt.Println("Input : ", input)
	fmt.Println("Answer: ", 4)
	fmt.Println("Output: ", p904.TotalFruit(input), "\n")

	input = []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
	fmt.Println("Input : ", input)
	fmt.Println("Answer: ", 5)
	fmt.Println("Output: ", p904.TotalFruit(input), "\n")
}

func test_p904t() {
	var input []int

	input = []int{1, 2, 3, 4, 4, 4, 3, 1, 5}
	fmt.Println("Input : ", input)
	fmt.Println("Answer: ", 5)
	fmt.Println("Output: ", p904.TotalFruit(input), "\n")
}
