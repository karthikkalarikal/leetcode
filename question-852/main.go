package main

func PeakIndexInMountainArray(arr []int) int {
	start, end := 0, len(arr)-1

	for start < end {
		mid := start + (end-start)/2

		if arr[mid] < arr[mid+1] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}
