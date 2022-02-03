package algorithms

func SelectionSort(nums []int) []int {
	numsLength := len(nums)
	for i := 0; i < numsLength; i++ {
		min := i
		for j := i + 1; j < numsLength; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
	return nums
}
