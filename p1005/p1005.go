//Given an array A of integers, we must modify the array in the following way:
// we choose an i and replace A[i] with -A[i],
// and we repeat this process K times in total.
// (We may choose the same index i multiple times.)
//
//Return the largest possible sum of the array after modifying it in this way.
//
//Example 1: Input: A = [4,2,3], K = 1
//Output: 5
//Explanation: Choose indices (1,) and A becomes [4,-2,3].
//Example 2: Input: A = [3,-1,0,2], K = 3
//Output: 6
//Explanation: Choose indices (1, 2, 2) and A becomes [3,1,0,2].
//Example 3: Input: A = [2,-3,-1,5,-4], K = 2
//Output: 13
//Explanation: Choose indices (1, 4) and A becomes [2,3,-1,5,4].
//
//Note:
//1 <= A.length <= 10000
//1 <= K <= 10000
//-100 <= A[i] <= 100
package main

import "fmt"

const zero int = 0

type mini struct {
	value int
	isNeg bool
}

func (min *mini) init(m int) {
	if m >= 0 {
		min.value, min.isNeg = m, false
	} else {
		min.value, min.isNeg = -m, true
	}
}

func (min *mini) setIfMin(m int) {
	if m < zero {
		if -m < min.value {
			min.value, min.isNeg = -m, true
		}
	} else {
		if m < min.value {
			min.value, min.isNeg = m, false
		}
	}
}

func largestSumAfterKNegations(A []int, K int) int {

	// 1. go through all number in A
	//    record the absolute of all negative ints
	//    record the int with minimum absolute value
	//    sum the non-neg
	var negInA []int
	var sum int
	// set the default minimum
	var min mini
	min.init(A[0])

	for _, a := range A {
		if a < zero {
			insertSort(&negInA, -a)
		} else {
			sum += a
		}
		min.setIfMin(a)
	}

	N := len(negInA)

	if N >= K {
		// 2.1 negate K largest negative value and get the sum
		sum += reverseAndSum(negInA, K)
		return sum
	} else { // K > N
		// 2.2 negate N largest negative value and get the sum
		sum += reverseAndSum(negInA, N)
		N = K - N
		if N%2 == 0 {
			return sum
		} else {
			return sum - 2*min.value
		}
	}
}

/*
 * pre: nums is sorted, higher index with higher value
 *      M <= N
 * the M largest value sum with positive sign
 * the rest sum with negative sign
 */
func reverseAndSum(nums []int, M int) (sum int) {
	N := len(nums)
	pos := nums[N-M:]
	neg := nums[:N-M]

	for i := range pos {
		sum += pos[i]
	}

	for j := range neg {
		sum -= neg[j]
	}

	return
}

func insertSort(nums *[]int, n int) {
	if len(*nums) == 0 {
		*nums = append(*nums, n)
		return
	}

	s := *nums
	for i, num := range s {
		if n > num {
			continue
		} else {
			tail := []int{n}
			tail = append(tail, s[i:]...)
			*nums = append(s[:i], tail...)
			return
		}
	}
	*nums = append(s, n)
	return
}

func main() {
	nums := []int{2, -3, -1, 5, -4}

	a := largestSumAfterKNegations(nums, 2)

	fmt.Println(nums, a)

}
