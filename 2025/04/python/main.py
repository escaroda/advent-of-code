from pathlib import Path


ROLL_CHAR = "@"
EMPTY_CHAR = "."
ROLLS_MAX = 3

dirs = [(-1, 0), (-1, 1), (0, 1), (1, 1), (1, 0), (1, -1), (0, -1), (-1, -1)]


def read_file(path: str) -> str:
    with open(path, encoding="utf-8") as f:
        return f.read()


def remove_rolls(grid: list[list[str]]) -> tuple[int, list[list[str]]]:
    rolls = 0
    grid_updated = [x[:] for x in grid]
    for i in range(len(grid)):
        for j in range(len(grid[i])):
            if grid[i][j] != ROLL_CHAR:
                continue

            count = 0
            for di, dj in dirs:
                if (
                    i + di < 0
                    or i + di >= len(grid)
                    or j + dj < 0
                    or j + dj >= len(grid[i])
                ):
                    continue

                if grid[i + di][j + dj] == ROLL_CHAR:
                    count += 1

            if count <= ROLLS_MAX:
                grid_updated[i][j] = EMPTY_CHAR
                rolls += 1

    return rolls, grid_updated


def part1(data: str) -> None:
    grid = list(map(list, data.splitlines()))
    count, _ = remove_rolls(grid)

    print("Part 1:", count)


def part2(data: str) -> None:
    grid = list(map(list, data.splitlines()))
    res = 0
    while True:
        count, grid = remove_rolls(grid)
        if count == 0:
            break

        res += count

    print("Part 2:", res)


def main():
    print("--- Day 4: Printing Department ---")
    print("https://adventofcode.com/2025/day/4\n")

    data = read_file(str(Path(__file__).parent.parent / "input.txt"))
    part1(data)
    part2(data)


if __name__ == "__main__":
    main()
