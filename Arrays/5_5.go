package chapter_5

func DeleteDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	writeIndex := 1
	for i := 1; i < len(nums); i++ {
		if nums[writeIndex-1] != nums[i] {
			nums[writeIndex] = nums[i]
			writeIndex++
		}
	}
	return writeIndex
}
