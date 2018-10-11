package main
import "fmt"

type INode struct {
	data string
	nextNode *INode
}
type QueueList struct {
	head *INode
	tail *INode
}

func main(){
	headNode:=INode{"头指针",nil}
	queue:=QueueList{&headNode,nil}
	addNode :=INode{"addNode",nil}
	headNode.nextNode = &addNode

	fmt.Println((*queue.head).data)
}
//入队操作
func enqueue(node INode,queue QueueList) bool{

}