/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    return add(l1, l2, 0)
}

func add (l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}

	sum := carry

	if l1 != nil {
		sum += l1.Val
		l1 = l1.Next
	}

	if l2 != nil {
		sum += l2.Val
		l2 = l2.Next
	}

	node := &ListNode {
		Val: sum%10,
	}

	node.Next = add(l1,l2, sum/10)

	return node
}
