package sort

import "fmt"

// 先大顶堆或者小顶堆
// 该顶就是排序要的第一个数字
// 将顶和底兑换位置，而后按照相同的逻辑重新堆排序 [0,len--]
// 重复以上操作
// 注：这边为了方便理解调试，append了一个0位后从1位开始计算
func HeapSort(list []int) {
	fmt.Println("heap sort:")
	newList := make([]int, 1)
	newList = append(newList, list...)
	heapify(&newList, len(newList), len(newList))
	for i := len(newList) - 1; i > 0; i-- {
		newList[1], newList[i] = newList[i], newList[1]
		heapify(&newList, i, i-1)
	}
	fmt.Println(newList[1:])
}

func heapify(list *[]int, idx, length int) {
	for i := len(*list) / 2; i > 0; i-- {
		heapAdjust(list, i, length)
	}
}
func heapAdjust(list *[]int, idx int, end int) {
	l := end
	li := 2 * idx
	ri := li + 1
	method := 0
	if l > li {
		if (*list)[idx] < (*list)[li] {
			method += 1 << 0
		}
	}
	if l > ri {
		if (*list)[idx] < (*list)[ri] {
			method += 1 << 1
		}
	}
	switch method {
	case 0:
		break
	case 1 << 0:
		(*list)[idx], (*list)[li] = (*list)[li], (*list)[idx]
		break
	case 1 << 1:
		(*list)[idx], (*list)[ri] = (*list)[ri], (*list)[idx]
		break
	default:
		if (*list)[li] > (*list)[ri] {
			(*list)[idx], (*list)[li] = (*list)[li], (*list)[idx]
		} else {
			(*list)[idx], (*list)[ri] = (*list)[ri], (*list)[idx]
		}
		break
	}

}
