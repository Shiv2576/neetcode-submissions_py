/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    head := l2
    carry := 0

    var prev *ListNode

    for l1 != nil || l2 != nil {
        sum := carry

        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }

        if l2 != nil {
            sum += l2.Val
            l2.Val = sum % 10
            carry = sum / 10

            prev = l2
            l2 = l2.Next
        } else {
            // l2 ended but l1 still exists
            prev.Next = &ListNode{
                Val: sum % 10,
            }
            carry = sum / 10
            prev = prev.Next
        }
    }

    if carry > 0 {
        prev.Next = &ListNode{Val: carry}
    }

    return head
}

