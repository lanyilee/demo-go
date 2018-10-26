package main

import "fmt"

//查找第一个等于value的下标
func BSearch1(A []int, low int, high int, value int) int {
	if A[low] > value || A[high] < value {
		return -1
	}
	for low <= high {
		mid := low + (high-low)>>1
		if A[mid] > value {
			high = mid - 1
		} else if A[mid] < value {
			low = mid + 1
		} else {
			if mid == 0 || A[mid-1] != value {
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

func main() {
	A := []int{1, 3, 4, 5, 7, 7, 8, 9, 10, 10, 10, 11}
	result := BSearch1(A, 0, len(A)-1, 7)
	fmt.Println(result)
}
