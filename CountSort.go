package main

import (
	"fmt"
	"strconv"
)

//计数排序
func CountSort(A []int, n int) []int {
	max := A[0]
	for _, a := range A {
		if a > max {
			max = a
		}
	}
	//创建新切片，长度为切片A的最大值加一
	C := make([]int, max+1, max+1) // 下标大小 [0,max]
	//根据下标统计各数
	for i := 0; i < n; i++ {
		C[A[i]]++
	}
	//C切片求和
	for i := 0; i <= max; i++ {
		if i > 0 {
			C[i] = C[i-1] + C[i]
		}
	}
	//下标不变，变的是下标对应的值（即个数）
	//排序
	R := make([]int, n, n) //储存排序好的数据
	//从高到低比较，A[i]看下它的前面有多少个人，从而决定自己的位置；稳定排序
	for i := n - 1; i >= 0; i-- {
		index := C[A[i]] - 1 //个数和减一，因为index为0开始
		R[index] = A[i]
		C[A[i]]--
	}
	A = R[:]
	return A
}

func main() {
	A := []int{3, 1, 2, 5, 2, 8, 4, 6, 9, 3, 0, 5, 3}
	A = CountSort(A, len(A))
	fmt.Println()
	for _, a := range A {
		s := strconv.Itoa(a) + ","
		fmt.Print(s)
	}

}
