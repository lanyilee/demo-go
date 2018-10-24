package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	//int_chan := make(chan int)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1 //如果int_chan为无缓冲，则主线程会堵塞,即此行之下代码都无法执行,除非有其他线程 <-int_chan,取走1
	string_chan <- "hello"
	fmt.Println(len(int_chan))
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}
