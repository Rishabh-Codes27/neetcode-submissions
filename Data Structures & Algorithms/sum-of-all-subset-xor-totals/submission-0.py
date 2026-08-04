class Solution:
    def subsetXORSum(self, nums: List[int]) -> int:
        def return_subsets(nums, index, current):
            
            if index == len(nums):
                return current
            
            #including current element
            including = return_subsets(nums, index+1, current ^ nums[index])

            #excluding current element
            excluding = return_subsets(nums, index+1, current)

            return including + excluding
        
        return return_subsets(nums, 0, 0)