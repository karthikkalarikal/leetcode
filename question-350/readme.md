# Intersection Of Two Arrays ii

## Problem Description

Given two integer arrays `nums1` and `nums2`, return an array of their intersection. Each element in the result must be __unique__ and you may return the result in __any order__.

> ### Follow up:
>> - What if the given array is already sorted? How would you optimize your algorithm?
>> - What if `nums1's` size is small compared to `nums2's` size? Which algorithm is better?
>> - What if elements of `nums2` are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?

- **Time Complexity:** O(n*log(n))
  - The time compexity is dominated by the sorting step. .

- **Space Complexity:** O(n)
  - The worst case would be that all of the elements are intersection ,hence the complexity.

## Solution

Sort two arrays, loop one using for loop then use each elecement as a target for the second array, since the memory of the second array is limited we will delete the element from that slice. Here i have taken the larger array as `num2`