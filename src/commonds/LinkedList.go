package commonds

import "log"

/**
链表结构
 */
type LinkedList struct {
	First *LinkNode
}

/**
创建 LinkedList
 */
func NewLinkedList() *LinkedList {
	return &LinkedList{nil}
}


/**
链表节点
 */
type LinkNode struct {
	Data interface{}
	Next *LinkNode
}

func (linkedlist *LinkedList) Insert(data interface{}) {
	newnode := &LinkNode{data,nil}
	if linkedlist.First == nil {
		linkedlist.First = newnode
		return
	}

	node := linkedlist.First
	for node.Next != nil {
		node = node.Next
	}
	node.Next = newnode
}

func (linkedlist *LinkedList) Operator(optr func(data interface{}) bool) {
	node := linkedlist.First
	for node != nil {
		if !optr(node.Data) {
			return
		}
		node = node.Next
	}
}

func (linkedlist *LinkedList) Show() {
	node := linkedlist.First
	for node != nil {
		log.Println(node.Data)
		node = node.Next
	}
}