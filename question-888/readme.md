# Fair Candy Swap

## Problem Statement

Alice and Bob have a different total number of candies. You are given two integer arrays `aliceSizes` and bobSizes where `aliceSizes[i]` is the number of candies of the `ith` box of candy that Alice has and `bobSizes[j]` is the number of candies of the jth box of candy that Bob has.

Since they are friends, they would like to exchange one candy box each so that after the exchange, they both have the same total amount of candy. The total amount of candy a person has is the sum of the number of candies in each box they have.

Return an integer array `answer` where `answer[0]` is the number of candies in the box that Alice must exchange, and `answer[1]` is the number of candies in the box that Bob must exchange. If there are multiple answers, you may return any one of them. It is guaranteed that at least one answer exists.

## Approach

The question can be done with a different algorithm more effectively, here the approch was to sort the arrays, take the total and find the difference, then iterate one of the arrys and find the one that will have half the difference in the other array.


- **Time Complexity:** O(n*log(n))
  - The time compexity is dominated by the sorting step.

- **Space Complexity:** O(log(n))
  - Helper function can stack so it would be O(1) + O(log n) which simplifies into O(log n)

