package sort

import "fmt"

//找最左边的值为参照值
//依次从两端的下标逐个往中间移动，碰到大于和小于的时候互相交换位置,最后将参照数和基准线交换位置
//然后基准线2端的分别再次用快排排序各自的小数组，直到排完
func QuickSort(list []int) {
	fmt.Println("quick sort:")
	newList := make([]int, 1)
	newList = append(newList, list...)
	quickSort(&newList, 1, len(newList))
	fmt.Println(newList[1:])
}

//10, 20, 50, 13, 106, 98, 77, 42, 91, 20
func quickSort(list *[]int, left, right int) {
	if left < right {
		middleIdx := partition(list, left, right)
		quickSort(list, left, middleIdx-1)
		quickSort(list, middleIdx+1, right)

	}
}

// a null b ...... c
// if b > a => null = b
// a b null ...... c
// if c > a => null = c
// a b c ...... null
//  -1 为了逻辑清楚
// 空位左右交替，不能连续出现在同一端
// 取一端为参照值
// 然后从另一端比较大小，逐个对比，满足条件后跳出
// 从另一个方向开始比较大小，逐个对比，满足条件后跳出
// 2端对比的不等号相反
func partition(list *[]int, start, end int) (middleIdx int) {
	reference := (*list)[start]
	(*list)[start] = -1
	(*list)[0] = reference
	location := start

	//plan1
	start++
	end--
	for start < end {
		for ; start <= end; end-- {
			if (*list)[end] < reference {
				(*list)[location], (*list)[end] = (*list)[end], -1
				location = end
				break
			}
		}
		for ; start <= end; start++ {
			if (*list)[start] > reference {
				(*list)[location], (*list)[start] = (*list)[start], -1
				location = start
				break
			}
		}

	}
	//plan2
	/*for i := end - 1; i > 0; i-- {
		if i == start {
			break
		}
		used := false
		for (*list)[i] < reference {
			(*list)[location], (*list)[i] = (*list)[i], -1
			location = i
			used = true
			break
		}
		if used {
			for j := start + 1; j < i; j++ {
				if (*list)[j] > reference {
					(*list)[location], (*list)[j] = (*list)[j], -1
					location = j
					start++
					break
				}
			}
		}
	}*/
	(*list)[location] = reference
	(*list)[0] = 0
	return location
}
