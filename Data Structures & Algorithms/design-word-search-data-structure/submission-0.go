type WordDictionary struct {
    children map[byte]*WordDictionary
    endof bool
}

func Constructor() *WordDictionary {
    return &WordDictionary {
        children : make(map[byte]*WordDictionary),
        endof : false,
    }
}

func (this *WordDictionary) AddWord(word string)  {
    cur := this
    for i := 0 ; i < len(word) ; i++ {
        ch := word[i]
        if cur.children[ch] == nil {
            cur.children[ch] = &WordDictionary {
                children : make(map[byte]*WordDictionary),
                endof : false,
            }
        }
        cur = cur.children[ch]
    }
    cur.endof = true
}

func (this *WordDictionary) Search(word string) bool {
    return this.searchHelper(word, 0) 
}


func (this *WordDictionary) searchHelper(word string, index int) bool {

    if index == len(word) {
        return this.endof
    }

    ch := word[index]

    if ch == '.' {
        for _ ,child := range this.children {
            if child != nil && child.searchHelper(word, index+1) {
                return true
            }
        }
        return false

    } else {
        if this.children[ch] == nil {
            return false 
        }
        return this.children[ch].searchHelper(word, index+1)
    }
}

