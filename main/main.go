// https://www.hackerrank.com/challenges/beautiful-triplets/problem

package main

import "fmt"

func contains(elements []int32, element1 int32, element2 int32, startPos int) bool {
	var foundCount = 0
	for i := startPos; i < len(elements); i++ {
		if foundCount == 0 && elements[i] == element1 {
			foundCount++
		} else if foundCount == 1 && elements[i] == element2 {
			foundCount++
			break
		}
	}

	return foundCount == 2
}

func beautifulTriplets(d int32, arr []int32) int32 {
	var count = int32(0)
	for idx, val := range arr {
		var nextVal = val + d
		var nextVal2 = val + 2*d

		if contains(arr, nextVal, nextVal2, idx) {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(beautifulTriplets(3, []int32{1, 2, 4, 5, 7, 8, 10,}))
}
