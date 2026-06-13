
func isPalindrome(s string) bool {

    cleaned := cleanstring(s)
    
    length := len(cleaned)

    for i := 0 ; i < length / 2 ; i++ {
        if cleaned[i] != cleaned[length-1-i] {
            return false
        }
    }

    return true

}


func cleanstring(s string) string {
    result := []byte{}
    
    for i := 0; i < len(s); i++ {
        ch := s[i]
        
        if ch >= 'a' && ch <= 'z' {
            result = append(result, ch)
        } else if ch >= 'A' && ch <= 'Z' {
            result = append(result, ch + 32) 
        } else if ch >= '0' && ch <= '9' {
            result = append(result, ch)
        }
    }
    
    return string(result)
}
