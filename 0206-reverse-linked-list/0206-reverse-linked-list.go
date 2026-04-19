/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    
	if head == nil {
		return head
	}

	var prev *ListNode = nil
	curr := head

	for curr != nil {
		next := curr.Next

		curr.Next = prev

		prev = curr
		curr = next
	}

	return prev
}