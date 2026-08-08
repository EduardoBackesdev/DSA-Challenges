package main

func merge(nums1 []int, m int, nums2 []int, n int) {

	i := m - 1
	j := n - 1
	l := len(nums1) - 1

	for j >= 0 {
		if i < 0 || nums2[j] > nums1[i] {
			nums1[l] = nums2[j]
			j--
		} else {
			nums1[l] = nums1[i]
			i--
		}

		l--
	}

}
