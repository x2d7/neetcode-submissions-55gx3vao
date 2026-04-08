class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        wanted: dict[int, int] = {}
        for i, num in enumerate(nums):
            want = target - num
            if num in wanted:
                return [wanted[num],i]
            wanted[want] = i
            