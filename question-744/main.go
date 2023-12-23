package main

func NextGreatestLetter(letters []byte, target byte) byte {
	end := len(letters) - 1
	first := 0

	nearest := 0
	if letters[end] < target {
		return letters[0]
	}

	for first <= end {
		mid := first + (end-first)/2
		// fmt.Printf("first : %v, end: %v , nearest: %v ", first, end, nearest)
		if target < letters[mid] {
			nearest = mid
			// fmt.Println("nearest", nearest)
			end = mid - 1
		} else {

			first = mid + 1
		}
		// nearest = mid
		// fmt.Printf("the result: %v expected: %v \n", string(letters[nearest]), string(target))
	}
	return letters[nearest]
}
