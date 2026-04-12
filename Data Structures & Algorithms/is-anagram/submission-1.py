class Solution:
    def isAnagram(self, s: str, t: str) -> bool:

        if len(s) != len(t):
            return False

        s_count = {}
        t_count = {}

        # for i in range(len(s)):
            
        #     s_count[s[i]] = s_count.get(s[i], 0) + 1
        #     t_count[t[i]] = t_count.get(t[i], 0) + 1

        # return 
 
        for x, y in zip(s, t):

            if x in s_count:
                s_count[x] += 1
            
            else:
                s_count[x] = 1

            if y in t_count:
                t_count[y] += 1

            else:
                t_count[y] = 1


        return s_count == t_count

            

            
        