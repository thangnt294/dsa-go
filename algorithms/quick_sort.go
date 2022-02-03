package algorithms

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition(nums, left, right)

	quickSort(nums, left, pivot-1)
	quickSort(nums, pivot+1, right)
}

func partition(nums []int, left, right int) int {
	pivot := right

	pointer := left

	for i := left; i < right; i++ {
		if nums[i] < nums[pivot] {
			nums[i], nums[pointer] = nums[pointer], nums[i]
			pointer++
		}
	}

	nums[pivot], nums[pointer] = nums[pointer], nums[pivot]
	return pointer
}
