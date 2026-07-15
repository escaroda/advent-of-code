from pathlib import Path


def read_file(path: str) -> str:
    with open(path, encoding="utf-8") as f:
        return f.read()


def part1(data: str) -> None:
    total = 0
    for line in data.splitlines():
        l, w, h = [int(x) for x in line.split("x")]
        side_a, side_b, side_c = l * w, w * h, h * l
        area = 2 * side_a+ 2 * side_b + 2 * side_c + min(side_a, side_b, side_c)
        total += area

    print("Part 1: ", total)


def part2(data: str) -> None:
    length = 0
    for line in data.splitlines():
        sides = [int(x) for x in line.split("x")]
        sides.sort()
        a, b, c = sides
        length += 2 * a + 2 * b + a * b * c
  
    print("Part 2: ", length)


def main():
    print("--- Day 2: I Was Told There Would Be No Math ---")
    print("https://adventofcode.com/2015/day/2\n")

    data = read_file(Path(__file__).parent.parent / "input.txt")
    part1(data)
    part2(data)


if __name__ == "__main__":
    main()
