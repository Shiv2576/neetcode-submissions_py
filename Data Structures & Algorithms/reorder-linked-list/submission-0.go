/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {

    mid := middle(head)

    second := mid.Next

    mid.Next = nil

    second = reverse(second)

    first := head

    for second != nil {
        tmp1 := first.Next
        tmp2 := second.Next

        first.Next = second
        second.Next = tmp1

        first = tmp1
        second = tmp2

    }
}


func reverse(head *ListNode) *ListNode {
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


func middle(head *ListNode) *ListNode {
    run1 := head
    run2 := head

    for run2 != nil && run2.Next != nil {
        run1 = run1.Next
        run2 = run2.Next.Next
    }

    return run1 
}