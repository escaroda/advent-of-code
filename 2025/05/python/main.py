from pathlib import Path
from typing import List


def read_file(path: str) -> str:
    with open(path, encoding="utf-8") as f:
        return f.read()


def parse_range(s: str) -> list[int]:
    return list(map(int, s.split("-")))


def part1(data: str) -> None:
    res = 0
    ranges_raw, ids_raw = data.strip().split("\n\n")
    ranges = list(map(parse_range, ranges_raw.splitlines()))
    ids = list(map(int, ids_raw.splitlines()))

    for id in ids:
        for l, r in ranges:
            if id < l or id > r:
                continue
            res += 1
            break

    print("Part 1:", res)


def merge_inclusive(ranges: List[List[int]]) -> List[List[int]]:
    ranges.sort()
    merged = [ranges[0]]

    for l, r in ranges[1:]:
        last_r = merged[-1][1]
        if l > last_r + 1:
            merged.append([l, r])
        else:
            merged[-1][1] = max(last_r, r)

    return merged


def part2(data: str) -> None:
    res = 0
    ranges_raw = data.strip().split("\n\n")[0]
    ranges = list(map(parse_range, ranges_raw.splitlines()))

    for l, r in merge_inclusive(ranges):
        res += r - l + 1

    print("Part 2:", res)


def main():
    print("--- Day 5: Cafeteria ---")
    print("https://adventofcode.com/2025/day/5\n")

    data = read_file(str(Path(__file__).parent.parent / "input.txt"))
    part1(data)
    part2(data)


if __name__ == "__main__":
    main()
