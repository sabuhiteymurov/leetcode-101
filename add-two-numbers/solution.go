package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func Solution(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	tail := dummy

	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := v1 + v2 + carry
		carry = sum / 10

		tail.Next = &ListNode{sum % 10, nil}
		tail = tail.Next
	}

	return dummy.Next
}
