from pathlib import Path
from typing import Iterator, Tuple


def parse(data: str) -> Iterator[Tuple[int, int, int]]:
    for line in data.splitlines():
        l, w, h = map(int, line.split("x"))
        yield l, w, h


def part1(data: str) -> int:
    total = 0
    for l, w, h in parse(data):
        side_a, side_b, side_c = l * w, w * h, h * l
        area = 2 * side_a + 2 * side_b + 2 * side_c + min(side_a, side_b, side_c)
        total += area

    return total


def part2(data: str) -> int:
    length = 0
    for dims in parse(data):
        a, b, c = sorted(dims)
        length += 2 * a + 2 * b + a * b * c

    return length


def main():
    print("--- Day 2: I Was Told There Would Be No Math ---")
    print("https://adventofcode.com/2015/day/2\n")

    file_path = Path(__file__).parent.parent / "input.txt"
    data = file_path.read_text(encoding="utf-8")
    print(f"Part 1: {part1(data)}")
    print(f"Part 2: {part2(data)}")


if __name__ == "__main__":
    main()
