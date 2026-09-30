class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        stack = []

        for num in tokens:
            if num == "+":
                stack[-2] = stack[-2] + stack[-1]
                stack.pop()

            elif num == "*":
                stack[-2] = stack[-2] * stack[-1]
                stack.pop()

            elif num == "-":
                stack[-2] = stack[-2] - stack[-1]
                stack.pop()

            elif num == "/":
                stack[-2] = int(stack[-2] / stack[-1])
                stack.pop()
            else:
                stack.append(int(num))
            
        

        return stack[-1]

            

        