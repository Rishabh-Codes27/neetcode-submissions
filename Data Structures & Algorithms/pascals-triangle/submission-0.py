class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        n = numRows
        res = [[1]]
        
        if n == 0:
            return []

        for _ in range(1, n):
            prev = res[-1]
            row = [1]
            for i in range(len(prev) - 1):
                row.append(prev[i] + prev[i + 1])
            row.append(1)
            res.append(row)
            
        return res