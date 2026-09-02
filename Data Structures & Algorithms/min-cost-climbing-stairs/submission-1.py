class Solution:
    def minCostClimbingStairs(self, cost: List[int]) -> int:
        n = len(cost)
        
        # dp[i] = minimum cost to reach step i
        dp = [0] * (n + 1)
        
        # You can start at step 0 or step 1 with 0 cost
        dp[0] = 0
        dp[1] = 0
        
        # For each step from 2 to n, calculate min cost to reach it
        for i in range(2, n + 1):
            # To reach step i, you come from either i-1 or i-2
            # and pay the cost of that previous step
            dp[i] = min(dp[i-1] + cost[i-1], dp[i-2] + cost[i-2])
        
        # The answer is the minimum cost to reach the top (step n)
        return dp[n]