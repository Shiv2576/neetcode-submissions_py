class Solution:
    def asteroidCollision(self, asteroids):
        stack = []

        for a in asteroids:
            alive = True

            while alive and stack and a < 0 and stack[-1] > 0:
                if stack[-1] < -a:
                    # Top asteroid explodes, current continues
                    stack.pop()
                elif stack[-1] == -a:
                    # Both explode
                    stack.pop()
                    alive = False
                else:
                    # Current asteroid explodes
                    alive = False

            if alive:
                stack.append(a)

        return stack