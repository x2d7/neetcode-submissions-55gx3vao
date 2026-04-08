from collections import Counter, defaultdict

class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        groups: dict[tuple, list[str]] = defaultdict(list)
        for s in strs:
            counter = Counter(s)
            key = tuple(sorted(counter.items()))
            groups[key].append(s)
        return list(groups.values())
