/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    runner1 := head
    runner2 := head


    for runner2 != nil && runner2.Next != nil {
        runner1 = runner1.Next
        runner2 = runner2.Next.Next

        if runner1 == runner2 {
            return true
        }
    }


    return false
    
}
