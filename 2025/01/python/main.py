from pathlib import Path


MIN_VAL = 0
MAX_VAL = 99
MOD_WRAP = MAX_VAL - MIN_VAL + 1
POS_START = 50


def read_file(path: str) -> str:
    with open(path, encoding="utf-8") as f:
        return f.read()


def part1(data: str) -> None:
    count = 0
    pos = POS_START
    for line in data.splitlines():
        delta = int(line[1:]) * (1 if line[0] == "R" else -1)
        pos = (pos + delta) % MOD_WRAP
        if pos == 0:
            count += 1

    print("Part 1: ", count)


def part2(data: str) -> None:
    count = 0
    pos = POS_START
    for line in data.splitlines():
        wraps, delta = divmod(int(line[1:]), MOD_WRAP)
        count += wraps
        delta *= 1 if line[0] == "R" else -1
        if pos + delta >= MOD_WRAP or (pos > 0 and pos + delta <= 0):
            count += 1
        pos = (pos + delta) % MOD_WRAP

    print("Part 2: ", count)


def main():
    print("--- Day 1: Secret Entrance ---")
    print("https://adventofcode.com/2025/day/1\n")

    data = read_file(str(Path(__file__).parent.parent / "input.txt"))
    part1(data)
    part2(data)


if __name__ == "__main__":
    main()
