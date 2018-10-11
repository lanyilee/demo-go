package main

import "fmt"

type INode struct {
	data     string
	nextNode *INode
}
type QueueList struct {
	head *INode
	tail *INode
}

func main() {
	headNode := INode{"头指针", nil}
	tailNode := INode{"尾指针", nil}
	queue := QueueList{&headNode, &tailNode}
	addNode := INode{"1", nil}
	addNode2 := INode{"2", nil}
	addNode3 := INode{"3", nil}
	addNode4 := INode{"4", nil}
	addNode5 := INode{"5", nil}
	enqueue(addNode, &queue)
	outQueue(&queue)
	enqueue(addNode2, &queue)
	enqueue(addNode3, &queue)
	enqueue(addNode4, &queue)
	enqueue(addNode5, &queue)
	outQueue(&queue)
	outQueue(&queue)
	outQueue(&queue)
	outQueue(&queue)
	outQueue(&queue)
}

//入队操作
func enqueue(node INode, queue *QueueList) bool {
	//队列无元素
	if queue.head.nextNode == queue.tail.nextNode && queue.head.nextNode == nil {
		queue.head.nextNode = &node
		queue.tail.nextNode = &node
	} else { //有多个元素
		var tempNode *INode
		tempNode = queue.head.nextNode
		for true {
			if tempNode.nextNode == nil {
				break
			}
			fmt.Println(tempNode.data)
			tempNode = tempNode.nextNode
		}
		(*tempNode).nextNode = &node
		queue.tail.nextNode = &node
	}
	return true
}

//出队操作
func outQueue(queue *QueueList) bool {
	//无元素
	if queue.head.nextNode == queue.tail.nextNode && queue.head.nextNode == nil {
		return false
	} else if queue.head.nextNode == queue.tail.nextNode { //只有一个元素
		queue.head.nextNode = queue.head.nextNode.nextNode
		queue.tail.nextNode = nil
	} else { //多个元素
		queue.head.nextNode = queue.head.nextNode.nextNode
	}
	return true
}
