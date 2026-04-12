class Solution:
    def isAnagram(self, s: str, t: str) -> bool:

        if len(s) != len(t):
            return False

        s_count = dict()
        t_count = dict()

        for x, y in zip(s, t):

            if x in s_count:
                s_count[x] += 1
            
            else:
                s_count[x] = 1

            if y in t_count:
                t_count[y] += 1

            else:
                t_count[y] = 1

        
        for char in s_count:
            if char not in t_count or s_count[char] != t_count[char]:
                return False

        return True

            

            
        