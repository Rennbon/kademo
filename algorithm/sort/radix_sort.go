package sort

import (
	"fmt"
	"math"
)

//先排序各位，在排序百位，在排序千位，以此类推
func RadixSort(list []int) {
	fmt.Println("radix sort:")
	max := list[0]
	for _, v := range list {
		if max < v {
			max = v
		}
	}
	log := int(math.Log10(float64(max)))
	for i := 1; i <= log; i++ {
		bucket(&list, int(math.Pow10(i)))
	}
	fmt.Println(list)
}

func bucket(list *[]int, e int) {
	bk := make([][]int, 10)
	capital := len(*list)/10 + 1
	for _, v := range *list {
		mod := v / e % 10
		if bk[mod] == nil {
			bk[mod] = make([]int, 0, capital)
		}
		bk[mod] = append(bk[mod], v)
	}
	i := 0
	for _, v := range bk {
		for _, vi := range v {
			(*list)[i] = vi
			i++
		}
	}
}
