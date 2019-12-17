package main

import "fmt"

var list = []int{10, 20, 50, 13, 106, 98, 77, 42, 91, 20}

func main() {
	//BubbleSort(list)
	//SelectionSort(list)
	//InsertionSort(list)
	QuickSort(list)
}

/*
冒泡排序
算法描述
- 比较相邻的元素。如果第一个比第二个大，就交换它们两个；
- 对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对，这样在最后的元素应该会是最大的数；
- 针对所有的元素重复以上的步骤，除了最后一个；
- 重复步骤1~3，直到排序完成。

前后对比，满足条件换位，
*/
func BubbleSort(list []int) {
	fmt.Println("bubble sort:")
	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	fmt.Println(list)
}

/* O(n2)
选择排序    每次从乱序中找到最小后放到新队列末端
算法描述
n个记录的直接选择排序可经过n-1次直接选择排序得到有序结果
- 初始状态：无序去为R[1,n],有序区为空
- 第i次排序（i=1,2,3,...,n-1）开始时，当前有序区和无序区分别为R[1,i-1]和R[i,n]。该次排序从当前无序区中选出关键字最小的记录
R[k],将它与无序区的第1个记录R交换，使得R[1,i]和R[i+1,n]分别变为记录个数增加1个的新有序区和记录个数减少1个的新无序区
- n-1次后，数组有序化
*/

func SelectionSort(list []int) {
	fmt.Println("selection sort:")
	count := len(list)
	for i := 0; i < count; i++ {
		min := list[i]
		for j := i + 1; j < count; j++ {
			if list[j] < min {
				min, list[j] = list[j], min
			}
		}
		list[i] = min
	}
	fmt.Println(list)
}

//从头开始遍历，每次去当前轮依次后前面的元素对比，如果满足条件，直接换位，碰到不满足则跳出
func InsertionSort(list []int) {
	fmt.Println("insertion sort:")
	//遍历多轮
	for i := 1; i < len(list); i++ {
		//没轮遍历已排序好的
		current := i
		for j := i; j > 0; j-- {
			//找到当前元素适合的位置
			if list[current] < list[j-1] {
				list[j-1], list[current] = list[current], list[j-1]
				current--
			} else {
				break
			}
		}
	}
	fmt.Println(list)
}

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

func partition2(arr []int, left, right int) int { // 分区操作
	pivot := left // 设定基准值（pivot）
	index := pivot + 1
	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {

			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}
	arr[pivot], arr[index-1] = arr[index-1], arr[pivot]
	return index - 1
}
