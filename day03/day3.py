"""
Advent of Code 2024
Day 3: Mull It Over

Usage:
    python3 day3.py part1 < input
    python3 day3.py part2 < input
"""

import argparse
import sys
import re

mul_regexp = re.compile(r"mul\((\d{1,3}),(\d{1,3})\)")


def main() -> None:
    parser = argparse.ArgumentParser()
    parser.add_argument("part", choices=["part1", "part2"])
    args = parser.parse_args()
    if args.part == "part1":
        exercise = sum_multiplicaton_results
    else:
        raise NotImplementedError()
    print(exercise(sys.stdin.read()))


def sum_multiplicaton_results(input: str) -> int:
    return sum(int(x) * int(y) for x, y in mul_regexp.findall(input))



if __name__ == "__main__":
    main()
