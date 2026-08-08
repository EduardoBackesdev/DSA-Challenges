package main

func removeElement(nums []int, val int) int {

	p1 := 0
	p2 := 1
	len := len(nums) - 1

	for p2 <= len {
		if nums[p1] == val {
			a := nums[p1]
			b := nums[p2]
			nums[p1] = b
			nums[p2] = a
			p2++
		}

	}

}
