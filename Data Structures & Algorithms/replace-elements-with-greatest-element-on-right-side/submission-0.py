class Solution:
    def replaceElements(self, arr: List[int]) -> List[int]:
        n = len(arr)
        last_digit = arr[n-1]
        result = [0] * n
        result[n-1] = -1

        for i in range(n-2 , -1 , -1):
            result[i] =  last_digit
            last_digit = max(last_digit , arr[i])
        

        return result

            

