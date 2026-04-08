from collections import Counter

class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        letters_s = Counter(s)
        letters_t = Counter(t)

        if letters_s != letters_t:
            return False
        return True