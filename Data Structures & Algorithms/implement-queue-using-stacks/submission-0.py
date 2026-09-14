class MyQueue:

    def __init__(self):
        self.stack_in = []   # for pushing new elements
        self.stack_out = []  # for popping/peeking elements

    def push(self, x: int) -> None:
        self.stack_in.append(x)

    def pop(self) -> int:
        self._transfer()
        return self.stack_out.pop()

    def peek(self) -> int:
        self._transfer()
        return self.stack_out[-1]

    def empty(self) -> bool:
        return not self.stack_in and not self.stack_out

    def _transfer(self):
        # Move all elements from stack_in to stack_out if stack_out is empty
        if not self.stack_out:
            while self.stack_in:
                self.stack_out.append(self.stack_in.pop())