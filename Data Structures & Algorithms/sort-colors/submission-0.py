class Solution:
    def sortColors(self, nums: List[int]) -> None:
        self.quickSort(nums , 0, len(nums) - 1)
    
    def quickSort(self , nums , low,  high):
        if low < high:
            pivot_index = self.partition(nums,low,high)
            self.quickSort(nums,low, pivot_index -1)
            self.quickSort(nums, pivot_index+1 , high  )

    
    def partition(self , nums, low, high):
        pivot = nums[low]
        i = low+1

        for j in range(low+1 , high+1):
            if nums[j] <= pivot:
                nums[i], nums[j] = nums[j] , nums[i]
                i += 1
        
        nums[low], nums[i - 1] = nums[i - 1], nums[low]
        return i - 1


        