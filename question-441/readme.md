# Arranging Coins

## Problem Description

You have n coins and you want to build a staircase with these coins. The staircase consists of k rows where the ith row has exactly i coins. The last row of the staircase may be incomplete.

## Solution

The solution is implemented in Go and uses an iterative approach to arrange the coins in rows. It keeps track of the remaining coins after each row and returns the total number of complete rows.

## Complexity Analysis
- **Time complexity**: O(n)
 - The function iterates through the rows linearly, making it linear in terms of the input n.
- **Space complexity**: O(1)
    - The algorithm uses a constant amount of space regardless of the input size. It doesn't require additional memory proportional to the input size.