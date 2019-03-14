package main

import (
	"fmt"
)

/* math block for continuous 1 multiplication and 1 division
 * return: N * (N-1) / (N-2)
 * return -1 for invalid attempt
 */
func multiplyAndDivide(N int) int {
	if N <= 0 {
		fmt.Println("ERROR: invalid input mAD()")
		return -1
	}
	if N <= 2 {
		return N
	}
	return N * (N - 1) / (N - 2)
}

/*
 * Second part of the equation
 * the value of the plus calculation
 * pV(k ,n) = k - 3
 * negative attempt are invalid
 */
func plusValue(k int) int {
	return k - 3
}

/*
 * Algorithm is
 * clumsy(N) = multiplyAndDivide(N) + plusValue(N) - mAD(N-4) + pV(N-4) - ...
 * the calculation ends when either of the function returns negative value
 */
func clumsy(N int) int {
	if N <= 0 {
		fmt.Println("ERROR: invalid input clumsy()")
		return -1
	}
	if N <= 3 {
		return multiplyAndDivide(N)
	}
	if N == 4 {
		return multiplyAndDivide(N) + plusValue(N)
	}

	ret := multiplyAndDivide(N) + plusValue(N)
	for n := N - 4; n > 0; n -= 4 {
		if m := multiplyAndDivide(n); m > 0 {
			ret -= m
		} else {
			break
		}

		if p := plusValue(n); p > 0 {
			ret += p
		} else {
			break
		}
	}
	return ret
}

func main() {

	for i := 1; i < 100; i++ {
		fmt.Println(clumsy(i), "-", i, "=", clumsy(i)-i)
	}
}
