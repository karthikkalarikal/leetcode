# CountNegatives Go Program

This Go program contains a function called `CountNegatives` designed to efficiently count negative numbers in a sorted 2D matrix.

## Function Description

To find out the number of negative numbers in a 2D array, the time complexity should be O(m+n)

## Solution Analysis 
- Using Binary Search in 2D array we can get the required time complexity, We find the upper and lower bound of the array and calculate the amount of -negative numbers, for instance since the array is sorted by row and collumn we find the upper bound when we find the -ve number and add the rest of the numbers in the column. 

## Complexity Analysis
- **Time complexity**: O(log(m+n))
 - The binary search algorithm reduces the search space in each step, resulting in a logarithmic time complexity.
- **Space complexity**: O(1)
    - The algorithm uses a constant amount of space regardless of the input size. It doesn't require additional memory proportional to the input size.