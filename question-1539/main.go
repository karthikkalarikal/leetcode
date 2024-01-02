package main

func FindKthPositive(arr []int, k int) int {
	start, end := 0, len(arr)-1
	mid := 0
	prev := 0
	diff := 0
	if arr[start] > k {
		return k
	}

	for start <= end {

		mid = start + (end-start)/2
		diff = arr[mid] - (mid + 1)

		if k <= diff {

			end = mid - 1

		} else {
			start = mid + 1

		}
	}

	prev = arr[start-1] + (k - (arr[start-1] - start))
	return prev

}
