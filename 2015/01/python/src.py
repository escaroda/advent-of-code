from pathlib import Path


def read_file(path: str) -> str:
    with open(path, encoding="utf-8") as f:
        return f.read()


def part1(data: str) -> None:
    floor = 0

    for ch in data:
        if ch == '(':
            floor += 1
        elif ch == ')':
            floor -= 1

    print("Part 1: ", floor)


def part2(data: str) -> None:
    floor = 0
    pos = 0

    for ch in data:
        pos += 1
        if ch == '(':
            floor += 1
        elif ch == ')':
            floor -= 1

        if floor == -1:
            break

    print("Part 2: ", pos)


def main():
    print("--- Day 1: Not Quite Lisp ---")
    print("https://adventofcode.com/2015/day/1\n")

    data = read_file(Path(__file__).parent.parent / "input.txt")
    part1(data)
    part2(data)


if __name__ == "__main__":
    main()
