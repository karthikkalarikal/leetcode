package main

/* Given a positive integer num, return true if num is a perfect square or false otherwise.

A perfect square is an integer that is the square of an integer. In other words, it is the product of some integer with itself.

You must not use any built-in library function, such as sqrt.
*/

func IsPerfectSquare(num int) bool {
	if num < 1 {
		return false
	}
	if num == 1 {
		return true
	}

	end := num / 2
	first := 2

	for end >= first {
		target := first + (end-first)/2 // find the middle , using this method to avoid index errors
		// fmt.Printf("target : %v first: %v end: %v num: %v \n", target, first, end, num)
		if target*target == num {
			return true
		} else if target*target < num {
			first = target + 1
		} else {
			end = target - 1
		}

	}
	return false
}
