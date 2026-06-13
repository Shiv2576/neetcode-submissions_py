/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {

    new := make(map[*Node]*Node)

    curr := head


    for curr != nil {
        new[curr] = &Node{Val : curr.Val}
        curr = curr.Next
    }


    curr = head

    for curr != nil {
        new[curr].Next = new[curr.Next]
        new[curr].Random = new[curr.Random]
        curr = curr.Next
    }

    return new[head]


    
}
