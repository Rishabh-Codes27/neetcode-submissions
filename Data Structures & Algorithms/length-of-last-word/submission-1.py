class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        tmp = []
        
        for i in range(len(s) -1, -1, -1):
            if s[i] == " ":
                if len(tmp) != 0:
                    break
                continue
            tmp.append(s[i])
        
        return len(tmp)
    
    