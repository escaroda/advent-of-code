from pathlib import Path
from itertools import combinations
from functools import reduce
from typing import Self
from pulp import LpProblem, LpStatus, LpMinimize, LpVariable, lpSum, LpInteger, PULP_CBC_CMD


class Machine:
    LIGHT_ON = "#"

    def __init__(
        self, lights: int, buttons: list[list[int]], joltage: list[int]
    ) -> None:
        self.lights = lights
        self.joltage = joltage
        self.buttons = buttons
        self.length = len(joltage)

    def btn_subsets(self) -> list[list[int]]:
        res = []
        merged = [self.pos_to_int(x, self.length) for x in self.buttons]
        for r in range(1, len(merged) + 1):
            for combo in combinations(merged, r):
                res.append(list(combo))

        return res
    
    def btns_as_pos(self) -> list[list[int]]:
        res = []
        for b in self.buttons:
            pos = [0] * self.length
            for i in b:
                pos[i] = 1
            res.append(pos)

        return res

    def press_btn(self, i: int, move: int = -1, times: int = 1) -> int:
        if move < 0:
            score = [0] * len(self.buttons[i])
            for j, pos in enumerate(self.buttons[i]):
                score[j] = self.joltage[pos]

            times = max(times, min(score))

        for pos in self.buttons[i]:
            self.joltage[pos] += move * times

        return times
    
    def unpress_btn(self, i: int, times: int = 1) -> None:
        self.press_btn(i, 1, times)

    @classmethod
    def from_str(cls, s: str) -> Self:
        lights_raw, s = s[1:].split("] (")
        buttons_raw, joltage_raw = s[:-1].split(") {")

        return cls(
            cls.bits_to_int([int(x == cls.LIGHT_ON) for x in lights_raw]),
            [[int(x) for x in b.split(",")] for b in buttons_raw.split(") (")],
            [int(x) for x in joltage_raw.split(",")],
        )

    @staticmethod
    def bits_to_int(bits: list[int]) -> int:
        res = 0
        for bit in bits:
            res = (res << 1) | bit
        return res

    @staticmethod
    def pos_to_int(pos: list[int], length: int) -> int:
        res = 0
        for p in pos:
            res |= 1 << (length - 1 - p)
        return res


def read_file(path: str) -> str:
    with open(path, encoding="utf-8") as f:
        return f.read()


def part1(data: str) -> None:
    res = 0
    machines = [Machine.from_str(m) for m in data.strip().splitlines()]
    for m in machines:
        for subset in m.btn_subsets():
            if m.lights == reduce(lambda a, b: a ^ b, subset):
                res += len(subset)
                break

    print("Part 1:", res)


def min_pressed(m: Machine) -> int:
    res = sum(m.joltage)

    def backtrack(count: int, btn_idx: int) -> None:
        nonlocal m
        nonlocal res
        if count >= res or any(x < 0 for x in m.joltage):
            return

        if sum(m.joltage) == 0:
            res = min(res, count)
            return

        for i in range(btn_idx, len(m.buttons)):
            times = m.press_btn(i)
            backtrack(count + times, i)
            m.unpress_btn(i, times)
            backtrack(count, i + 1)

    backtrack(0, 0)
    return res


def part2(data: str) -> None:
    res = 0
    machines = [Machine.from_str(m) for m in data.strip().splitlines()]
    for m in machines:
        buttons = m.btns_as_pos()

        prob = LpProblem()
        x = [LpVariable(f"x{i}", lowBound=0, cat=LpInteger) for i in range(len(buttons))]

        for j in range(len(m.joltage)):
            prob += lpSum(x[i] * buttons[i][j] for i in range(len(buttons))) == m.joltage[j]

        prob += lpSum(x)
        status = prob.solve(PULP_CBC_CMD(msg=False))
        if LpStatus[status] != "Optimal":
            raise ValueError("No optimal solution")

        res += sum(int(xi.value() or 0) for xi in x)

    print("Part 2:", res)


def main():
    print("--- Day 10: Factory ---")
    print("https://adventofcode.com/2025/day/10\n")

    data = read_file(str(Path(__file__).parent.parent / "input.txt"))
    part1(data)
    part2(data)


if __name__ == "__main__":
    main()
