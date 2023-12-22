package main

/* You have n coins and you want to build a staircase with these coins. The staircase consists of k rows where the ith row has exactly i coins. The last row of the staircase may be incomplete.

Given the integer n, return the number of complete rows of the staircase you will build.
*/

func ArrangeCoins(n int) int {
	target := n
	for i := 1; i <= n; i++ {
		target -= i
		if target < 0 {
			return i - 1
		}
		// fmt.Printf("num: %v, i: %v \n", n, i)
	}
	return 1
}
