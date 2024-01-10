package main

import (
	"fmt"
	"sort"
)

func FairCandySwap(aliceSizes []int, bobSizes []int) []int {
	sort.Ints(aliceSizes)
	sort.Ints(bobSizes)
	aliceTot := 0
	bobTot := 0
	ansArray := make([]int, 2)
	for _, v := range aliceSizes {
		aliceTot += v
	}

	for _, v := range bobSizes {
		bobTot += v
	}

	diff := aliceTot - bobTot

	for _, v := range aliceSizes {
		ansArray = candyHelper(bobSizes, v, -diff)

		if ansArray[0] != -1 {
			break
		}
		// fmt.Println(ansArray)
	}
	return ansArray

}

func candyHelper(arr1 []int, v, target int) []int {

	start, end := 0, len(arr1)-1

	for start <= end {
		mid := start + (end-start)/2
		// fmt.Println("a,b,t", v, arr1[mid], target)
		fmt.Println("target,arr1[mid]", target, arr1[mid]-v)
		if 2*(arr1[mid]-v) == target {
			return []int{v, arr1[mid]}
		} else if 2*(arr1[mid]-v) < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return []int{-1, -1}
}

