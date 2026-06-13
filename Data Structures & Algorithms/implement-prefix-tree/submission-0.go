type PrefixTree struct {
    children map[byte]*PrefixTree
    endofWord bool
}

func Constructor() *PrefixTree {
    return &PrefixTree {
        children : make(map[byte]*PrefixTree),
        endofWord : false,
    }
}

func (this *PrefixTree) Insert(word string) {
    cur := this

    for i := 0 ; i < len(word) ; i++ {
        ch := word[i]
        if cur.children[ch] == nil {
            cur.children[ch] = &PrefixTree {
                children : make(map[byte]*PrefixTree),
                endofWord :  false,
            }
        }
        cur = cur.children[ch]
    }
    cur.endofWord = true
}

func (this *PrefixTree) Search(word string) bool {
    cur := this

    for i := 0 ; i < len(word) ; i++ {
        ch := word[i]
        if cur.children[ch] == nil {
            return false
        }
        cur = cur.children[ch]
    }

    return cur.endofWord
}

func (this *PrefixTree) StartsWith(prefix string) bool {
    cur := this

    for i := 0 ; i < len(prefix) ; i ++ {
        ch := prefix[i]
        if cur.children[ch] == nil {
            return false
        }

        cur = cur.children[ch]
    }

    return true
}