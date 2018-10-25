package main

import (
	"fmt"
)

func Qvot(num float64) float64 {
	//num = math.Trunc(num*1e6 + 0.5)*1e-6
	var root float64
	if num <= 0 {
		return 0
	} else if num == 1 {
		return 1
	} else if num < 1 {
		root = GetRootOne(num, 1.0, num)
	} else if num > 1 {
		root = GetRootOne(1.0, num, num)
	}
	return root
}

func GetRootOne(small float64, big float64, num float64) float64 {
	mid := (big + small) / 2
	//mid = math.Trunc(mid*1e1 + 0.5)*1e-1
	fil := mid * mid
	if fil < num-0.0000001 {
		mid = GetRootOne(mid, big, num)
	} else if fil > num+0.0000001 {
		mid = GetRootOne(small, mid, num)
	} else if fil > num-0.0000010 || fil < num+0.0000001 {
		fmt.Println(mid)
	}
	return mid
}

func main() {
	x := 90.0
	Qvot(x)
}
