package algorithms

func InsertionSort(nums []int) []int {
	numsLength := len(nums)
	for i := 1; i < numsLength; i++ {
		key := nums[i]
		j := i - 1
		for j >= 0 && key < nums[j] {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}
	return nums
}
