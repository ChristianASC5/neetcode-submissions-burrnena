class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:

        i = 0
        j = 1

        result = nums[i] + nums[j]

        while result != target:

            if j < len(nums) - 1:
                j += 1

            else:
                i += 1
                j = i + 1
            
            result = nums[i] + nums[j]
            

        return [i, j]


        # 2, 3, 4, 6, 8