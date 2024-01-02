# SearchInsert Algorithm

## Description

The `SearchInsert` function is designed to solve the problem of finding the index of a target value in a sorted array or, if the target is not present, return the index where it would be inserted to maintain the sorted order. The algorithm has a runtime complexity of O(log n), making it efficient for large datasets.

## Problem Statement

Given a sorted array of distinct integers and a target value, the function aims to:

- Return the index if the target is found in the array.
- If the target is not found, return the index where it would be inserted to maintain the sorted order.


- **Time Complexity:** O(log n)
  - The algorithm utilizes a binary search approach, efficiently reducing the search space in each iteration.

- **Space Complexity:** O(1)
  - The algorithm uses a constant amount of extra space, regardless of the size of the input array. The space required is independent of the array size.


## Function Signature

```go
func SearchInsert(nums []int, target int) int