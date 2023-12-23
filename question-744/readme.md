# Next Greatest Letter

## Problem Description

You are given an array of characters
letters that is sorted in non-decreasing order,
and a character target.

## Solution

The solution is implemented in Go and uses a binary search algorithm to efficiently find the desired character in the sorted array.

## Complexity Analysis
- **Time complexity**: O(log(n))
 - The binary search algorithm reduces the search space in each step, resulting in a logarithmic time complexity.
- **Space complexity**: O(1)
    - The algorithm uses a constant amount of space regardless of the input size. It doesn't require additional memory proportional to the input size.