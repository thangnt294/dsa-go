package algorithms

func BubbleSort(nums []int) []int {
	numsLength := len(nums)
	for i := 0; i < numsLength; i++ {
		for j := 0; j < numsLength-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
