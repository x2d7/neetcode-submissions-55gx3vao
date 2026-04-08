from collections import Counter
from typing import List


class TopK:
    def __init__(self, k: int, key=None, reverse=False):
        self._k = k
        self._arr = []
        self._key = key
        self._reverse = reverse

    def add(self, n):
        self._arr.append(n)
        self._arr.sort(key=self._key, reverse=self._reverse)
        self._arr = self._arr[:self._k]

    def get(self) -> list:
        return self._arr


class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        freqs = list(enumerate(Counter(nums).items()))
        max_test = lambda x: x[1][1]
        get_num = lambda x: x[1][0]
        topk = TopK(k, reverse=True, key=max_test)

        for f in freqs:
            topk.add(f)

        result = topk.get()
        return [get_num(x) for x in result]