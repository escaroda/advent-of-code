import operator
from pathlib import Path
from functools import reduce


EMPTY_CHAR = " "
ADD_CHAR = "+"


def read_file(path: str) -> str:
    with open(path, encoding="utf-8") as f:
        return f.read()


def part1(data: str) -> None:
    res = 0
    lines = data.splitlines()
    rows = [x.strip().split() for x in lines[:-1]]
    cols = [list(map(int, row)) for row in zip(*rows)]  # rotated rows
    ops = [
        operator.add if x == ADD_CHAR else operator.mul
        for x in lines[-1].strip().split()
    ]

    for i, nums in enumerate(cols):
        res += reduce(ops[i], nums)

    print("Part 1:", res)


def part2(data: str) -> None:
    res = 0
    lines = data.splitlines()
    digit_length = len(lines) - 1
    ops_raw = lines[-1]
    edges = [i for i, v in enumerate(ops_raw) if v != EMPTY_CHAR] + [len(ops_raw) + 1]

    ops = []
    for i in range(len(edges) - 1):
        op = operator.add if ops_raw[edges[i]] == ADD_CHAR else operator.mul
        ops.append((op, edges[i], edges[i + 1] - 1))

    for op, start, end in ops:
        digits = []
        for i in range(start, end):
            digit = []
            for row in range(digit_length):
                if lines[row][i] != EMPTY_CHAR:
                    digit.append(lines[row][i])
            digits.append(int("".join(digit)))

        res += reduce(op, digits)

    print("Part 2:", res)


def main():
    print("--- Day 6: Trash Compactor ---")
    print("https://adventofcode.com/2025/day/6\n")

    data = read_file(str(Path(__file__).parent.parent / "input.txt"))
    part1(data)
    part2(data)


if __name__ == "__main__":
    main()
