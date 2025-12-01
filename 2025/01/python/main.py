import operator
from pathlib import Path


MIN_VAL = 0
MAX_VAL = 99
MOD_WRAP = MAX_VAL - MIN_VAL + 1
POS_START = 50

ops = {
    'R': operator.add,
    'L': operator.sub,
}


def read_file(path: str) -> str:
    with open(path, encoding="utf-8") as f:
        return f.read()


def part1(data: str) -> None:
    count = 0
    pos = POS_START
    for line in data.replace("L", "-").replace("R", "").splitlines():
        pos += int(line)
        pos %= MOD_WRAP
        if pos == 0:
            count += 1

    print("Part 1: ", count)


def part2(data: str) -> None:
    count = 0
    pos = POS_START
    for line in data.splitlines():
        for _ in range(int(line[1:])):
            op = ops[line[0]]
            pos = op(pos, 1)
            pos %= MOD_WRAP
            if pos == 0:
                count += 1

    print("Part 2: ", count)


def main():
    print("--- Day 1: Secret Entrance ---")
    print("https://adventofcode.com/2025/day/1\n")

    data = read_file(str(Path(__file__).parent.parent / "input.txt"))
    part1(data)
    part2(data)


if __name__ == "__main__":
    main()
