from pathlib import Path


BEAM_ENTER = "S"
SPLITTER = "^"


def read_file(path: str) -> str:
    with open(path, encoding="utf-8") as f:
        return f.read()


def part1(data: str) -> None:
    res = 0
    lines = data.splitlines()
    beams = {lines[0].find(BEAM_ENTER)}
    lvl = 1
    while lvl < len(lines):
        beams_next: set[int] = set()
        for i in beams:
            if lines[lvl][i] == SPLITTER:
                res += 1
                beams_next.add(i - 1)
                beams_next.add(i + 1)
            else:
                beams_next.add(i)

        beams = beams_next
        lvl += 1

    print("Part 1:", res)


def part2(data: str) -> None:
    lines = data.splitlines()
    w = len(lines[0])
    beams = [0] * w
    beams[lines[0].find(BEAM_ENTER)] = 1
    lvl = 1
    while lvl < len(lines):
        beams_next = [0] * w
        for i, v in enumerate(beams):
            if lines[lvl][i] == SPLITTER:
                beams_next[i - 1] += v
                beams_next[i + 1] += v
            else:
                beams_next[i] += v

        beams = beams_next
        lvl += 1

    print("Part 2:", sum(beams))


def main():
    print("--- Day 7: Laboratories ---")
    print("https://adventofcode.com/2025/day/7\n")

    data = read_file(str(Path(__file__).parent.parent / "input.txt"))
    part1(data)
    part2(data)



if __name__ == "__main__":
    main()
