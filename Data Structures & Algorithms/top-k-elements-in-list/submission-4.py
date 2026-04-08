from collections import Counter
from typing import List
import bisect

class TopK:
    def __init__(self, k: int):
        self._k = k
        self._arr: List[tuple] = []

    def add(self, index_num_freq: tuple):
        index, (num, freq) = index_num_freq
        entry = (-freq, index, num)
        bisect.insort_right(self._arr, entry)
        if len(self._arr) > self._k:
            self._arr.pop()

    def get(self) -> list:
        return [t[2] for t in self._arr]

class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        freqs = list(enumerate(Counter(nums).items()))
        topk = TopK(k)
        for f in freqs:
            topk.add(f)
        return topk.get()
