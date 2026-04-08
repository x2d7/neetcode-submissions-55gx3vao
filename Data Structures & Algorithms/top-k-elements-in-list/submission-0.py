from collections import Counter

MIN_INT = -1001

class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        result = []
        freqs = list(enumerate(Counter(nums).items()))
        max_test = lambda x: x[1][1] if x is not None else MIN_INT

        for _ in range(k):
            top = max(freqs, key=max_test)
            i, (num, freq) = top
            freqs[i] = None
            result.append(num)

        return result
