package main

import "fmt"

//删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow

}
func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
