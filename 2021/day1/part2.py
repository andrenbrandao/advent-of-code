import sys


def depths_increased(depths):
    count = 0

    for i in range(3, len(depths)):
        if (
            depths[i - 3] + depths[i - 2] + depths[i - 1]
            < depths[i - 2] + depths[i - 1] + depths[i]
        ):
            count += 1
    return count


if __name__ == "__main__":
    input_file = sys.argv[1]

    lines = []
    with open(input_file) as file:
        lines = [int(line.strip()) for line in file]

    print(depths_increased(lines))
