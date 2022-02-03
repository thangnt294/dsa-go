package main

import (
	"dsa-go/algorithms"
	"fmt"
	"time"
)

func main() {
	nums := []int{1, 2, 7, 14, 42, 22, 9, 15, 21, 24, 19, 8, 0, 28, 99, 45, 67, 42, 32, 51, 613, 344, 623, 542, 242, 111, 54, 65635, 6234, 54, 345, 333, 485, 87,
  423, 424, 6564, 3, 63, 22222, 484, 302, 5, 120, 1022, 3938, 1222, 948, 11, 102, 23, 77, 9281, 233}
  bubble := make([]int, len(nums));
  selection := make([]int, len(nums));
  insertion := make([]int, len(nums));
  merge := make([]int, len(nums));
  quick := make([]int, len(nums));

  copy(bubble, nums)
  copy(selection, nums)
  copy(insertion, nums)
  copy(merge, nums)
  copy(quick, nums)

  for i := 1; i <= 10; i++ {
    start0 := time.Now()
    algorithms.BubbleSort(bubble)
    t0 := time.Since(start0)
    start1 := time.Now()
    algorithms.SelectionSort(selection)
    t1 := time.Since(start1)
    start2 := time.Now()
    algorithms.InsertionSort(insertion)
    t2 := time.Since(start2)
    start3 := time.Now()
    algorithms.MergeSort(merge)
    t3 := time.Since(start3)
    start4 := time.Now()
    algorithms.QuickSort(quick)
    t4 := time.Since(start4)
    fmt.Printf("Bubble: %v | Selection: %v | Insertion: %v | Merge: %v | Quick: %v\n", t0, t1 , t2, t3, t4);
    time.Sleep(time.Second)
  }
}
