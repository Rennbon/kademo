package sort

import "fmt"

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
