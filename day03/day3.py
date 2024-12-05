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
from collections.abc import Generator

mul_regexp = re.compile(r"mul\((\d{1,3}),(\d{1,3})\)")
condition_regexp = re.compile(r"do(?:n't)?\(\)")


def main() -> None:
    parser = argparse.ArgumentParser()
    parser.add_argument("part", choices=["part1", "part2"])
    args = parser.parse_args()
    if args.part == "part1":
        exercise = sum_multiplicaton_results
    else:
        exercise = sum_conditional_multiplication_results
    print(exercise(sys.stdin.read()))


def sum_multiplicaton_results(input: str) -> int:
    return sum(int(x) * int(y) for x, y in mul_regexp.findall(input))


def sum_conditional_multiplication_results(input: str) -> int:
    return sum(
        sum_multiplicaton_results(subpart) for subpart in find_all_enabled(input)
    )


def find_all_enabled(input: str) -> Generator[str, None, None]:
    start = 0
    for match in condition_regexp.finditer(input):
        if match.group() == "do()":
            if start is not None:
                # Only the most recent do() instruction applies.
                continue
            start = match.end()
        elif start is not None:  # match.group() == "don't()" and do() was found
            yield input[start:match.start()]
            start = None
    if start is not None:
        yield input[start:]


if __name__ == "__main__":
    main()
