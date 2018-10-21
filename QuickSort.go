package main

import (
	"fmt"
	"strconv"
)

func QuickSort(A []int, p int, r int) {
	if p >= r {
		return
	}
	q := Partition(A, p, r)
	QuickSort(A, p, q-1)
	QuickSort(A, q+1, r)

}

func Partition(A []int, p int, r int) int {
	key := A[r]
	i := p
	j := p
	for j < r {
		if A[j] < key { //遍历的A[j]数小于key就交换,保证A[i]的左边小于key，右边大于key
			tem := A[i]
			A[i] = A[j]
			A[j] = tem
			i++
		}
		j++
	}
	A[r] = A[i]
	A[i] = key
	return i
}

func main() {
	A := []int{3, 1, 21, 5, 2, 8, 4, 6, 9, 13, 10, 15, 7}
	QuickSort(A, 0, len(A)-1)
	fmt.Println()
	for _, a := range A {
		s := strconv.Itoa(a) + ","
		fmt.Print(s)
	}
}
