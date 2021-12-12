import sys
from typing import ForwardRef


class Submarine:
    def __init__(self, depth=0, horizontal_position=0):
        self.depth = depth
        self.horizontal_position = horizontal_position

    def execute(self, action, value):
        if action == "forward":
            self.forward(value)
        elif action == "down":
            self.down(value)
        elif action == "up":
            self.up(value)
        else:
            raise Exception("Invalid action")

    def forward(self, value):
        self.horizontal_position += value

    def down(self, value):
        self.depth += value

    def up(self, value):
        self.depth -= value


def pilot_submarine(commands):
    submarine = Submarine()

    for command in commands:
        action, value = command.split()
        submarine.execute(action, int(value))

    return submarine.depth, submarine.horizontal_position


if __name__ == "__main__":
    input_file = sys.argv[1]

    lines = []
    with open(input_file) as file:
        lines = [line.strip() for line in file]

    depth, horizontal_position = pilot_submarine(lines)
    print(depth * horizontal_position)
