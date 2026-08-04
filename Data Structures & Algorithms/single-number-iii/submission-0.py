class Solution:
    def singleNumber(self, nums: List[int]) -> List[int]:
        n, res = len(nums), []

        for i in range(n):
            flag = True
            for j in range(n):
                if i != j and nums[i] == nums[j]:
                    flag = False
                    break

            if flag:
                res.append(nums[i])
                if len(res) == 2:
                    break

        return res
