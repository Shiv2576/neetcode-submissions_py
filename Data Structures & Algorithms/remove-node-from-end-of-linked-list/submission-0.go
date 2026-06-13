/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // Edge case: empty list
    if head == nil {
        return nil
    }
    
    // Calculate length
    length := 0
    curr := head
    for curr != nil {
        length++
        curr = curr.Next
    }
    
    // If n is greater than length or invalid
    if n <= 0 || n > length {
        return head
    }
    
    // Special case: removing the first node (nth from end = first from start)
    if n == length {
        return head.Next
    }
    
    // Find the node to remove (position from start)
    element := length - n
    
    // Find and remove the node
    matched := 0
    current := head
    var prev *ListNode
    
    for current != nil {
        if matched == element {
            // Remove the node
            if prev == nil {
                // This should not happen due to above check, but just in case
                head = current.Next
            } else {
                prev.Next = current.Next
            }
            break
        }
        
        prev = current
        current = current.Next
        matched++
    }
    
    return head
}