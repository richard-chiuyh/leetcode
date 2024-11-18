/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	head := new(ListNode)
	carry := 0
	current := head
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		current.Val = sum % 10
		carry = sum / 10
		if l1 != nil || l2 != nil || carry != 0 {
			current.Next = new(ListNode)
			current = current.Next
		}
	}
	return head
}

// @lc code=end

