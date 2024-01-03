package main

func CountNegatives(grid [][]int) int {
	negativeNumbers := 0

	r, c := 0, len(grid[0])-1
	// fmt.Println("c", c)
	// fmt.Println("len(grid) ", len(grid))
	for r < len(grid) && c >= 0 {
		// fmt.Println("c", c)
		if grid[r][c] < 0 {
			negativeNumbers += len(grid) - r
			// fmt.Printf("len(grid) %d -r %d = %d\n", len(grid), r, negativeNumbers)

		}
		// fmt.Println("num ", grid[r][c])
		if grid[r][c] >= 0 {
			r++
		} else {
			c--

		}
		// fmt.Println("c, r ,negative", c, r, negativeNumbers)
	}

	return negativeNumbers
}
