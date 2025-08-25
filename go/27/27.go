package main

import "fmt"

func removeElement(nums []int, val int) int {
	var count int
	var sliceDif []int
	lenghtInitial := len(nums)
	for idx, value := range nums {
		if value == val {
			nums[idx] = 0
			count++
			continue
		}
		sliceDif = append(sliceDif, value)
	}
	copy(nums, sliceDif)
	return lenghtInitial - count
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	result := removeElement(nums, val)
	fmt.Println(result)
	fmt.Println(nums)
}
