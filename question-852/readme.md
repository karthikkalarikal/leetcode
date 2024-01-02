# Peak Index In Mountain Array

## Problem Statement

An array `arr` is considered a mountain if the following properties hold:

- `arr.length >= 3`
- There exists some `i` with `0 < i < arr.length - 1` such that:
  - `arr[0] < arr[1] < ... < arr[i - 1] < arr[i]`
  - `arr[i] > arr[i + 1] > ... > arr[arr.length - 1]`

The task is to find and return the index `i` such that `arr[0] < arr[1] < ... < arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1]`.

Additionally, the solution must have a time complexity of O(log(arr.length)).

## Approach

The given problem can be efficiently solved using binary search. The basic idea is to narrow down the search space based on the comparison of adjacent elements.

The implemented Go function `PeakIndexInMountainArray` takes an array `arr` as input and returns the index `i` satisfying the mountain array conditions. It uses a binary search approach to achieve O(log(arr.length)) time complexity.

Here's how the algorithm works:

1. Initialize two pointers, `start` and `end`, representing the range of the array to be considered.
2. Perform a binary search within this range until `start` becomes equal to `end`.
3. At each step, calculate the middle index `mid`.
4. Compare `arr[mid]` with `arr[mid+1]`.
   - If `arr[mid]` is less than `arr[mid+1]`, it means the peak lies to the right. So, update `start = mid + 1`.
   - Otherwise, the peak lies to the left. Update `end = mid`.
5. Repeat steps 3-4 until `start` becomes equal to `end`.
6. Return the `start` index, as it represents the peak index in the mountain array.

## Usage

To use this function, simply call it with a mountain array as an argument:

```go
arr := []int{...} // Input mountain array
result := PeakIndexInMountainArray(arr)
fmt.Println("Peak Index:", result)
