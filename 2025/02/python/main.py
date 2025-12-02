from pathlib import Path


def read_file(path: str) -> str:
    with open(path, encoding="utf-8") as f:
        return f.read()
    

def parse_range(s: str) -> list[int]:
    return list(map(int, s.split("-")))


def int_split(n: int, mod=10) -> list[int]:
    res = []
    while n:
        res.append(n % mod)
        n //= mod
    return res


def invalid1(i: int) -> bool:
    digits = int_split(i)
    l = len(digits)
    if l % 2 == 1:
        return False
    
    mid = l // 2
    for i in range(mid):
        if digits[i] != digits[mid + i]:
            return False
    
    return True


def part1(data: str) -> None:
    res = 0
    for part in data.split(","):
        r = parse_range(part)
        for n in range(r[0], r[1]+1):
            if invalid1(n):
                res += n

    print("Part 1: ", res)


def invalid2(n: int) -> bool:
    l = len(int_split(n))
    for i in range(1, l // 2 + 1):
        if l % i > 0:
            continue

        digits = int_split(n, 10**i)
        unique = set(digits)
        if len(unique) == 1:
            return True
        
    return False


def part2(data: str) -> None:
    res = 0
    for part in data.split(","):
        r = parse_range(part)
        for n in range(r[0], r[1]+1):
            if invalid2(n):
                res += n

    print("Part 2: ", res)


def main():
    print("--- Day 2: Gift Shop ---")
    print("https://adventofcode.com/2025/day/2\n")

    data = read_file(str(Path(__file__).parent.parent / "input.txt"))
    part1(data)
    part2(data)


if __name__ == "__main__":
    main()
