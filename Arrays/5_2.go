package chapter_5

func IncrementAnInteger(nums []int) []int {
	nums[len(nums)-1]++
	for i := len(nums) - 1; i > 0 && nums[i] == 10; i-- {
		nums[i] = 0
		nums[i-1]++
	}
	if nums[0] == 10 {
		nums[0] = 1
		nums = append(nums, 0)
	}
	return nums
}
