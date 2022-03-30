package main

import "fmt"

var list = []int{50, 10, 20, 1239, 13, 106, 98, 77, 42, 91, 20}

func main() {
	/*sort.BubbleSort(list)
	sort.SelectionSort(list)
	sort.InsertionSort(list)

	sort.HeapSort(list)
	sort.MergeSort(list)*/
	//sort.CountingSort(list)
	//sort.RadixSort(list)

	//sort.QuickSort(list)
	//Qsort1(list)
	M1(list)
}

func M1(arr []int) {
	res := Merge(arr)
	for _, v := range res {
		fmt.Println(v)
	}
}
func Merge(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	middle := (len(arr) + 1) / 2
	res := Merge2(Merge(arr[:middle]), Merge(arr[middle:]))
	return res
}
func Merge2(arr1, arr2 []int) []int {
	l := len(arr1) + len(arr2)
	res := make([]int, 0, l)

	l1 := 0
	l2 := 0
	for {
		p := 0
		if l1 >= len(arr1) {
			p += 1
		}
		if l2 >= len(arr2) {
			p += 2
		}
		b := false
		switch p {
		case 0:
			if arr1[l1] > arr2[l2] {
				res = append(res, arr2[l2])
				l2++
			} else {
				res = append(res, arr1[l1])
				l1++
			}
			continue
		case 1:
			res = append(res, arr2[l2])
			l2++
			continue
		case 2:
			res = append(res, arr1[l1])
			l1++
		default:
			b = true
			break
		}
		if b {
			break
		}
	}
	return res
}

func Qsort1(list []int) {
	Qsort2(list, 0, len(list)-1)
	for _, v := range list {
		fmt.Println(v)
	}
}

func Qsort2(arr []int, left, right int) {
	if left < right {
		mid := partitoin(arr, left, right)
		Qsort2(arr, left, mid)
		Qsort2(arr, mid+1, right)
	}
}
func partitoin(arr []int, left, right int) int {
	cmp := left
	idx := cmp + 1
	for i := idx; i <= right; i++ {
		if arr[i] < arr[cmp] {
			arr[i], arr[idx] = arr[idx], arr[i]
			idx++
		}
	}
	arr[cmp], arr[idx-1] = arr[idx-1], arr[cmp]
	return idx - 1
}
