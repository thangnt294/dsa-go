package algorithms

func MergeSort(nums []int) []int {
	left := 0;
	right := len(nums) - 1;
	return mergeSort(nums, left, right);
}

func mergeSort(nums []int, left, right int) []int {
	if (left != right) {
		middle := (left + right) / 2;
		firstArray := mergeSort(nums, left, middle);
		secondArray := mergeSort(nums, middle + 1, right);

		// merge the 2 arrays
		mergeArray := []int{};
		firstIndex := 0;
		secondIndex := 0;
		firstArrayLength := len(firstArray);
		secondArrayLength := len(secondArray);
		
		for firstIndex < firstArrayLength && secondIndex < secondArrayLength {
			// merge base on the smallest element of the 2 arrays
			var element int;
			if firstArray[firstIndex] < secondArray[secondIndex] {
				element = firstArray[firstIndex];
				firstIndex++;
			} else {
				element = secondArray[secondIndex];
				secondIndex++;
			}
			mergeArray = append(mergeArray, element);
		}

		for firstIndex < firstArrayLength {
			mergeArray = append(mergeArray, firstArray[firstIndex]);
			firstIndex++;
		}

		for secondIndex < secondArrayLength {
			mergeArray = append(mergeArray, secondArray[secondIndex]);
			secondIndex++;
		}

		return mergeArray;
	}
	return []int{nums[left]};
}