package sort

import "fmt"

//数据量大，数据值分布量小
//取最大值最小值之间的长度做桶
//然后在指定桶对应处自增
func CountingSort(list []int) {
	fmt.Println("counting sort:")
	l := len(list)
	ret := make([]int, 0, l)
	min, max := list[0], list[0]
	for i := 1; i < l; i++ {
		if min > list[i] {
			min = list[i]
		}
		if max < list[i] {
			max = list[i]
		}
	}
	bucket := make([]int, max-min+1)

	for i := 0; i < l; i++ {
		bucket[list[i]-min]++
	}
	for i := 0; i < len(bucket); i++ {
		for bucket[i] > 0 {
			ret = append(ret, i+min)
			bucket[i]--
		}
	}
	fmt.Println(ret)
}
