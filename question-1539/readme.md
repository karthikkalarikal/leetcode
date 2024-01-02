# FindKthPositive Algorithm

## Time and Space Complexity

The `FindKthPositive` algorithm has the following complexities:

- **Time Complexity:** O(log n)
  - The algorithm employs a binary search strategy to efficiently locate the Kth missing positive integer in the array.

- **Space Complexity:** O(1)
  - The algorithm utilizes a constant amount of extra space, irrespective of the array's size. The space required is not influenced by the input array size.

## Function Signature

```go
func FindKthPositive(arr []int, k int) int
