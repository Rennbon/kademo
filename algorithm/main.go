package main

import "github.com/Rennbon/kademo/algorithm/sort"

var list = []int{10, 20, 50, 1239, 13, 106, 98, 77, 42, 91, 20}

func main() {
	/*sort.BubbleSort(list)
	sort.SelectionSort(list)
	sort.InsertionSort(list)
	sort.QuickSort(list)
	sort.HeapSort(list)
	sort.MergeSort(list)*/
	//sort.CountingSort(list)
	sort.RadixSort(list)
}
