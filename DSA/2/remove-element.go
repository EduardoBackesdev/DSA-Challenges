package main

func removeElement(nums []int, val int) int {

	p1 := 0
	p2 := len(nums) - 1

	for p1 <= p2 {
		if nums[p1] == val {
			a := nums[p1]
			b := nums[p2]
			nums[p1] = b
			nums[p2] = a
			p2--
		} else {
			p1++
		}
	}

	return p1

}
