package merge

import (
	"slices"
)

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	if n == 0 {
		return nums1
	}
	nums1 = nums1[:m]
	nums1 = append(nums1, nums2...)
	slices.Sort(nums1)
	return nums1

}
func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 6, nums2, 3)
}
