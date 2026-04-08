class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        appears: dict[int, bool] = {}
        for num in nums:
            if num in appears:
                return True
            appears[num] = True
        return False