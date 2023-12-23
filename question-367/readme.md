# Is Perfect Square

## Problem Description

Given a positive integer `num`, return true if `num` is a perfect square or false otherwise. A perfect square is an integer that is the square of another integer. In other words, it is the product of some integer with itself.

You must not use any built-in library function, such as `sqrt`.

## Solution

The solution is implemented in Go and uses a binary search algorithm to efficiently determine whether a given positive integer is a perfect square.

## Complexity Analysis
- **Time complexity**: O(log(n))
 - The binary search algorithm reduces the search space in each step, resulting in a logarithmic time complexity.
- **Space complexity**: O(1)
    - The algorithm uses a constant amount of space regardless of the input size. It doesn't require additional memory proportional to the input size.