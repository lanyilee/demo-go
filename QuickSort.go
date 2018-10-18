package main

func QuickSort(A []*int, p int, r int) {
	if p >= r {
		return
	}
	q := Partition(A, p, r)
	QuickSort(A, p, q-1)
	QuickSort(A, q+1, r)

}

func Partition(A []*int, p int, r int) int {

}

func main() {

}
