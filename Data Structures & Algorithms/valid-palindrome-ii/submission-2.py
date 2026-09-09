class Solution:
    def validPalindrome(self, s: str) -> bool:
        
        left , right = 0 , len(s) - 1
                
        def valid(left , right) -> bool:
            while left < right:
                if s[left] != s[right]:
                    return False
                
                left +=1
                right -=1

            return True

        while left < right:
            if s[left] != s[right]:
                return valid(left + 1, right) or valid(left, right - 1)

            left += 1
            right -= 1


        return True

