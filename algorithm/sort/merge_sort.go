package sort

import "fmt"

//将一个数组分解到单len数组，然后两两合并合并在合并
//每个合并的都会变成有序的，然后依次对比大小添加到有序数组的下一位
func MergeSort(list []int) {
	fmt.Println("merge sort:")
	arr := mergeSort(list)
	fmt.Println(arr)
}

func mergeSort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	l := (len(list) + 1) / 2
	ls := list[:l]
	lr := list[l:]
	return merge(mergeSort(ls), mergeSort(lr))
}

func merge(arr1, arr2 []int) []int {
	cp := len(arr1) + len(arr2)
	//存放有序新数组
	ret := make([]int, 0, cp)
	i, j := 0, 0
	for i < len(arr1) || j < len(arr2) {
		f := 0
		if len(arr1) <= i {
			f = 1
		}
		if len(arr2) <= j {
			f = 2
		}
		switch f {
		case 0:
			if arr1[i] < arr2[j] {
				ret = append(ret, arr1[i])
				i++

			} else {
				ret = append(ret, arr2[j])
				j++
			}
		case 1:
			ret = append(ret, arr2[j])
			j++

		case 2:
			ret = append(ret, arr1[i])
			i++
		}
	}
	return ret
}
