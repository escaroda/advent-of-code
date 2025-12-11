from pathlib import Path
from functools import lru_cache


YOU_NODE = "you"
SVR_NODE = "svr"
DAC_NODE = "dac"
FFT_NODE = "fft"
TARGET_NODE = "out"


def read_file(path: str) -> str:
    with open(path, encoding="utf-8") as f:
        return f.read()


def graph_from_str(data: str) -> dict[str, list[str]]:
    graph: dict[str, list[str]] = {}
    for line in data.strip().splitlines():
        device, line = line.split(": ")
        attachments = line.split(" ")
        graph[device] = attachments

    return graph


def part1(data: str) -> None:
    graph = graph_from_str(data)
    count = 0

    def dfs(node: str, visited: set[str]) -> None:
        nonlocal count
        if node == TARGET_NODE:
            count += 1
            return

        for n in graph.get(node, []):
            if n not in visited:
                visited.add(n)
                dfs(n, visited)
                visited.remove(n)

    dfs(YOU_NODE, set(YOU_NODE))

    print("Part 1:", count)


def part2(data: str) -> None:
    graph = graph_from_str(data)

    @lru_cache(maxsize=None)
    def count_paths(node: str, dac_visited: bool, fft_visited: bool) -> int:
        if node == TARGET_NODE:
            return int(dac_visited and fft_visited)

        dac_visited = dac_visited or node == DAC_NODE
        fft_visited = fft_visited or node == FFT_NODE

        total = 0
        for n in graph.get(node, []):
            total += count_paths(n, dac_visited, fft_visited)

        return total

    print("Part 2:", count_paths(SVR_NODE, False, False))


def main():
    print("--- Day 11: Reactor ---")
    print("https://adventofcode.com/2025/day/11\n")

    data = read_file(str(Path(__file__).parent.parent / "input.txt"))
    part1(data)
    part2(data)


if __name__ == "__main__":
    main()
