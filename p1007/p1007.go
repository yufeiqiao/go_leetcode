package main

import "fmt"

func minDominoRotations(A []int, B []int) int {
	aA, aB, bA, bB := 0, 1, 1, 0
	for i := 1; i < len(A); i++ {
		if A[i] == B[0] {

		}
	}
	return aA
}

func main() {
	A := []int{2, 1, 2, 4, 2, 2}
	B := []int{5, 2, 6, 2, 3, 2}

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println("Output: ", minDominoRotations(A, B))
}
