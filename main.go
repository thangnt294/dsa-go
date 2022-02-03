package main

import (
	"dsa-go/algorithms"
	"fmt"
)

func main() {
	nums := []int{1, 2, 7, 14, 42, 22, 9, 15}
	algorithms.InsertionSort(nums)
	fmt.Println(nums)
}
