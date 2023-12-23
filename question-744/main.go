package main

/*

You are given an array of characters
letters that is sorted in non-decreasing order,
and a character target.
There are at least two different characters in letters.
Return the smallest character in letters that is
lexicographically greater than target.
 If such a character does not exist,
 return the first character in letters.

*/

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
