import heapq
import math
from pathlib import Path
from itertools import combinations


PAIRS_LIMIT = 1000  # 10 for test input
N_LARGEST_CIRCUITS = 3


class UnionFind:
    def __init__(self) -> None:
        self.parent: dict[int, int] = {}

    def find(self, x: int)-> int:
        if x not in self.parent:
            self.parent[x] = x
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]

    def union(self, a: int, b: int) -> None:
        ra, rb = self.find(a), self.find(b)
        if ra != rb:
            self.parent[rb] = ra

    def groups(self) -> list[list[int]]:
        result: dict[int, list[int]] = {}
        for x in self.parent:
            root = self.find(x)
            result.setdefault(root, []).append(x)
        return list(result.values())


def read_file(path: str) -> str:
    with open(path, encoding="utf-8") as f:
        return f.read()


def part1(data: str) -> None:
    boxes = [
        tuple(int(x) for x in line.split(",")) for line in data.strip().splitlines()
    ]
    distances = [
        ((math.dist(boxes[i], boxes[j]), i, j))
        for i, j in combinations(range(len(boxes)), 2)
    ]
    distances.sort()
    uf = UnionFind()
    for _, i, j in distances[:PAIRS_LIMIT]:
        uf.union(i, j)

    sizes = [len(g) for g in uf.groups()]
    largest_circuits = heapq.nlargest(N_LARGEST_CIRCUITS, sizes)

    print("Part 1:", math.prod(largest_circuits))


def part2(data: str) -> None:
    res = 0
    boxes = [
        tuple(int(x) for x in line.split(",")) for line in data.strip().splitlines()
    ]
    distances = [
        ((math.dist(boxes[i], boxes[j]), i, j))
        for i, j in combinations(range(len(boxes)), 2)
    ]
    distances.sort()
    uf = UnionFind()
    processed: set[int] = set()
    for _, i, j in distances:
        uf.union(i, j)
        processed.add(i)
        processed.add(j)
        if len(processed) == len(boxes) and len(uf.groups()) == 1:
            res = boxes[i][0] * boxes[j][0]
            break

    print("Part 2:", res)


def main():
    print("--- Day 8: Playground ---")
    print("https://adventofcode.com/2025/day/8\n")

    data = read_file(str(Path(__file__).parent.parent / "input.txt"))
    part1(data)
    part2(data)


if __name__ == "__main__":
    main()
