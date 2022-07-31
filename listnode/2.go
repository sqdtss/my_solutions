package solution

//2. 两数相加
//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//请你将两个数相加，并以相同形式返回一个表示和的链表。
//你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//示例 1：
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
//
//示例 2：
//输入：l1 = [0], l2 = [0]
//输出：[0]
//
//示例 3：
//输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//输出：[8,9,9,9,0,0,0,1]
//
//提示：
//每个链表中的节点数在范围 [1, 100] 内
//0 <= Node.val <= 9
//题目数据保证列表表示的数字不含前导零

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	pre := &ListNode{}
//	cur := pre
//	flag := 0
//	for l1 != nil || l2 != nil {
//		l1Val, l2Val := 0, 0
//		if l1 != nil {
//			l1Val = l1.Val
//		}
//		if l2 != nil {
//			l2Val = l2.Val
//		}
//		val := l1Val + l2Val + flag
//		flag = 0
//		cur.Next = &ListNode{val, nil}
//		cur = cur.Next
//		if cur.Val >= 10 {
//			cur.Val %= 10
//			flag = 1
//		}
//		if l1 != nil {
//			l1 = l1.Next
//		}
//		if l2 != nil {
//			l2 = l2.Next
//		}
//	}
//	if flag == 1 {
//		cur.Next = &ListNode{1, nil}
//	}
//	return pre.Next
//}
