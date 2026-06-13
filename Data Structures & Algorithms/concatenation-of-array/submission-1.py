class Solution:
    def getConcatenation(self, nums: List[int]) -> List[int]:
        n = len(nums)
        result = [0] * 2*n

        for i in range(2*n):
            if i < n:
                result[i] = nums[i]
            else:
                result[i] = nums[i-n]

        return result
