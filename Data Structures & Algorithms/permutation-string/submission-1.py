class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:

        if len(s1) > len(s2):
            return False
        

        s1_count = {}

        for letter in s1:
            s1_count[letter] = s1_count.get(letter , 0) + 1
            
        

        window = len(s1)

        for i in range(len(s2) - window + 1):
            window_count = s2[i : i+window]

            window_1 = {}

            for letter in window_count:
                window_1[letter] = window_1.get(letter , 0) + 1

            if window_1 == s1_count:
                return True
        

        return False
        