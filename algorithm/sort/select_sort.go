package sort

import "fmt"

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
