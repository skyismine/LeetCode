package problems

/**
题目描述：
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

https://leetcode-cn.com/problems/add-two-numbers/
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	switch METHOD {
	case METHOD_MINE:
		return addTwoNumbersMine(l1, l2)
	case METHOD_LCA:
		return addTwoNumbersLCA(l1, l2)
	}

	return nil
}

/**
思路：遍历l1和l2将每项相加并处理进位情况
复杂度分析：
时间复杂度: O(n)，n为l1和l2中较长的链表长度
空间复杂度：O(n), n为l1和l2中较长的链表长度
 */
func addTwoNumbersMine(l1 *ListNode, l2 *ListNode) *ListNode {
	first := l1
	second := l2
	//相加后是否进位
	overflow := 0
	var listnode,ret *ListNode = nil,nil
	for first != nil || second != nil {
		val := overflow
		// 一次性处理两个链表的判空问题，不为空就处理
		if first != nil {
			val += first.Val
			first = first.Next
		}
		if second != nil {
			val += second.Val
			second = second.Next
		}
		//进位处理
		if val >= 10 {
			val -= 10
			overflow = 1
		} else {
			overflow = 0
		}
		//插入链表
		if listnode == nil {
			listnode = &ListNode{val, nil}
			ret = listnode
		} else {
			listnode.Next = &ListNode{val, nil}
			listnode = listnode.Next
		}
	}
	//遍历完成后如果overflow仍为1说明最高位进1需要多插入一个节点
	if overflow == 1 {
		listnode.Next = &ListNode{overflow, nil}
	}
	return ret
}

/**
LCA解法和我的思路相同
 */
func addTwoNumbersLCA(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersMine(l1, l2)
}

/**
拓展

如果链表中的数字不是按逆序存储的呢？例如：

(3 -> 4 -> 2) + (4 -> 6 -> 5) = 8 -> 0 -> 7

思路：可以使用递归调用的方式先算低位后算高位，然后使用和上面相同的思路处理进位
 */