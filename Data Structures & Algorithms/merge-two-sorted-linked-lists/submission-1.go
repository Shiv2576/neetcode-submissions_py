/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}

    tail := dummy

    curr1 := list1
    curr2 := list2

    for curr1 != nil && curr2 != nil {
        if curr1.Val <= curr2.Val {
            tail.Next = curr1
            curr1 = curr1.Next
        } else {
            tail.Next = curr2
            curr2 = curr2.Next
        }

        tail = tail.Next
    }

    if curr1 != nil {
        tail.Next = curr1
    } else {
        tail.Next = curr2
    }


    return dummy.Next
}

