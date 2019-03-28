package commonds

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
