package main

func Intersection(nums1 []int, nums2 []int) []int {

	m := []int{}
	mp := make(map[int]bool)

	for _, v := range nums1 {
		for _, w := range nums2 {
			if v == w && !mp[v] {
				mp[v] = true
				m = append(m, v)
				continue
			}
		}
	}
	return m
}
