package main

import "fmt"

//异常处理机制
func main() {
	RecoverHandle()
	fmt.Println("can go")
}

func RecoverHandle() {
	//MyPanic() 放到这里没效，defer一定得在可能异常前
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	MyPanic()
}

func MyPanic() {
	a := 10
	b := 0
	y := a / b
	fmt.Print(y)
}
