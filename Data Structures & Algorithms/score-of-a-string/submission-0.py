class Solution:
    def scoreOfString(self, s: str) -> int:
        n = len(s)
        sum = 0
        for i in range(n-1):
            value1 = ord(s[i])
            value2 = ord(s[i+1])

            sum += abs(value1 - value2)


        return sum


