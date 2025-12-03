from pathlib import Path


def read_file(path: str) -> str:
    with open(path, encoding="utf-8") as f:
        return f.read()


def part1(data: str) -> None:
    res = 0
    for line in data.splitlines():
        digits = list(map(int, list(line)))

        l_max = (1, 0)  # (value, index)
        r_max = (1, len(digits) - 1)

        for i in range(l_max[1], r_max[1]):
            if digits[i] > l_max[0]:
                l_max = (digits[i], i)

        for i in range(r_max[1], l_max[1], -1):
            if digits[i] > r_max[0]:
                r_max = (digits[i], i)

        res += l_max[0] * 10 + r_max[0]

    print("Part 1:", res)


def part2(data: str) -> None:
    batteries = 12
    res = 0
    for line in data.splitlines():
        digits = list(map(int, list(line)))
        joltage = 0
        left_border = 0
        for b in range(1, batteries + 1):
            max_value = 0
            for i in range(left_border, len(digits) - batteries + b):
                if digits[i] > max_value:
                    max_value = digits[i]
                    left_border = i + 1

            joltage += max_value * 10 ** (batteries - b)

        res += joltage

    print("Part 2:", res)


def main():
    print("--- Day 3: Lobby ---")
    print("https://adventofcode.com/2025/day/3\n")

    data = read_file(str(Path(__file__).parent.parent / "input.txt"))
    part1(data)
    part2(data)


if __name__ == "__main__":
    main()
