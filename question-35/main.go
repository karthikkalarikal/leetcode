package main

func SearchInsert(nums []int, target int) int {
	first, end := 0, len(nums)-1
	mid := 0

	if target < nums[0] {
		return 0
	}

	for first <= end {
		mid = first + (end-first)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			first = mid + 1
		}
	}
	// fmt.Printf("target: %v, index: %v\n", target, mid)
	if target > nums[mid] {
		return mid + 1
	} else {
		return mid
	}

}
