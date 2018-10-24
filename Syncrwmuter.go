package main

import (
	"fmt"
	"strconv"
	"sync"
)

type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	//ch := make(chan interface{})
	ch := make(chan interface{}, 3)
	//fmt.Println(len(ch))
	//ch<-1
	//fmt.Println(len(ch))
	//ch<-1
	//fmt.Println(len(ch))
	//ch<-1
	//fmt.Println(len(ch))
	go func() {
		set.RLock()
		for elem, value := range set.s {
			ch <- value
			fmt.Println("elem:" + strconv.Itoa(elem))
			//开始这里困惑为什么len(ch)在两次循环后才是1，原因是第一次流入之后，被外面主线程“v := <-th.Iter() ”,流出去了
			fmt.Println(len(ch))
		}
		close(ch)
		set.RUnlock()
	}()
	return ch
}

func main() {

	th := threadSafeSet{
		s: []interface{}{"1", "2", "3"},
	}
	//这里会堵塞，等待th.Iter()流进数据，之后才流出到v
	v := <-th.Iter()
	fmt.Println("over")
	fmt.Println(v)
}
