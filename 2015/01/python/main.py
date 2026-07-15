from pathlib import Path


def part1(data: str) -> int:
    floor = 0

    for ch in data:
        if ch == "(":
            floor += 1
        elif ch == ")":
            floor -= 1

    return floor


def part2(data: str) -> int:
    floor = 0
    pos = 0

    for ch in data:
        pos += 1
        if ch == "(":
            floor += 1
        elif ch == ")":
            floor -= 1

        if floor == -1:
            break

    return pos


def main():
    print("--- Day 1: Not Quite Lisp ---")
    print("https://adventofcode.com/2015/day/1\n")

    file_path = Path(__file__).parent.parent / "input.txt"
    data = file_path.read_text(encoding="utf-8")
    print(f"Part 1: {part1(data)}")
    print(f"Part 2: {part2(data)}")


if __name__ == "__main__":
    main()
