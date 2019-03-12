package main

import "fmt"

/*
 * 1002. Find Common Characters
 * Easy
 * Given an array A of strings made only from lowercase letters,
 * return a list of all characters that show up in all strings within the list
 * (including duplicates).
 * For example, if a character occurs 3 times in all strings but not 4 times,
 * you need to include that character three times in the final answer.
 * You may return the answer in any order.
 *
 * Example 1: Input: ["bella","label","roller"]
 * Output: ["e","l","l"]
 *
 * Example 2: Input: ["cool","lock","cook"]
 * Output: ["c","o"]
 * Note:
 * 1 <= A.length <= 100
 * 1 <= A[i].length <= 100
 * A[i][j] is a lowercase letter
 */

func initMap(s string) map[rune]int {
	ret := make(map[rune]int)
	for _, r := range []rune(s) {
		if _, has := ret[r]; has {
			ret[r]++
		} else {
			ret[r] = 1
		}
	}
	return ret
}

func updateMap(rM map[rune]int, s string) (map[rune]int, []rune) {
	ret := make(map[rune]int)
	content := []rune{}
	for _, r := range []rune(s) {
		if _, has := rM[r]; has {
			ret[r]++
			content = append(content, r)
			if rM[r] == 1 {
				delete(rM, r)
			} else {
				rM[r]--
			}
		}
	}
	return ret, content
}

func commonChars(A []string) []string {
	if len(A) <= 1 {
		return A
	}

	rMap := initMap(A[0])
	content := []rune{}

	for _, s := range A[1:] {
		rMap, content = updateMap(rMap, s)
	}

	return toStringSlice(content)
}

func toStringSlice(runes []rune) []string {
	strings := []string{}
	for _, r := range runes {
		strings = append(strings, string(r))
	}
	return strings
}

func main() {
	s := "strings"
	r := "stri"

	array := commonChars([]string{s, r})

	fmt.Println(s, r, array)
}
