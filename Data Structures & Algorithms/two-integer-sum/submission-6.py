class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:

        nums_dict = { val: index for index, val in enumerate(nums)}

        for i in range(len(nums)):

            diff = target - nums[i]

            if diff in nums_dict:

                if i != nums_dict[diff]:

                    return [i, nums_dict[diff]]

        return []


        # 2, 3, 4, 6, 8