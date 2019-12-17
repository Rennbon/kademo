package main

import "github.com/Rennbon/kademo/algorithm/sort"

var list = []int{10, 20, 50, 13, 106, 98, 77, 42, 91, 20}

func main() {
	//BubbleSort(list)
	//SelectionSort(list)
	//InsertionSort(list)
	//QuickSort(list)
	sort.HeapSort(list)
}
