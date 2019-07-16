package problems

/**
问题描述：
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

https://leetcode-cn.com/problems/merge-two-sorted-lists/
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	switch METHOD {
	case METHOD_MINE:
		return mergeTwoListsMine(l1, l2)
	case METHOD_LCA:
		return mergeTwoListsLCA(l1, l2)
	}

	return nil
}

/**
算法：
先查找第一个值比较小的链表，以它为基准进行遍历，设第一个链表为fl1，第二个链表为fl2，如果fl1遍历完了fl2还没完则把fl2剩余节点附加到fl1上
遍历时比较如果fl1的当前值小于fl2且fl1下个值大于fl2则说明fl2应该插入到当前位置，需要将fl1.Next设置为fl2、fl2.Next设置为原本fl1.Next，
同时将fl2置为原始fl2.Next准备下一轮比较
否则将fl1下移一个，进行下一轮比较

时间复杂度: O(n)，n为fl1的长度
空间复杂度: o(1)
 */
func mergeTwoListsMine(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	fl1 := l1
	fl2 := l2
	if (l1.Val > l2.Val) {
		fl1 = l2
		fl2 = l1
	}
	firstnode := fl1
	for fl1 != nil && fl2 != nil {
		node1 := fl1.Next
		node2 := fl2.Next
		if node1 == nil {
			fl1.Next = fl2
			break
		} else if fl1.Val <= fl2.Val && node1.Val > fl2.Val {
			fl1.Next = fl2
			fl2.Next = node1
			fl2 = node2
		}
		fl1 = fl1.Next
	}

	return firstnode
}

/**
LCA上暂无解答
 */
func mergeTwoListsLCA(l1 *ListNode, l2 *ListNode) *ListNode {
	return nil
}