class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        seen = {}

        for i, x in enumerate(nums):
            if x in seen:
                return True
            else:
                seen[x] = i

        return False

        