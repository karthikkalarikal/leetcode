package main

import (
	"sort"
)

func Intersection(nums1 []int, nums2 []int) []int {
	lnum1 := len(nums1)
	lnum2 := len(nums2)

	sort.Ints(nums1)
	sort.Ints(nums2)
	ans := make([]int, 2)
	if lnum1 <= lnum2 {
		ans = intersectionHelper(nums1, nums2)
	} else {
		ans = intersectionHelper(nums2, nums1)
	}
	return ans
}

func intersectionHelper(num1 []int, num2 []int) []int {
	m := []int{}
	

	for _, v := range num1 {	
		start, end := 0, len(num2)-1
		for start <= end {

			mid := start + (end-start)/2

			if num2[mid] == v {
				
				m = append(m, num2[mid])
				num2 = append(num2[:mid], num2[mid+1:]...)
				break
			} else if num2[mid] > v {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}

	}
	return m
}
