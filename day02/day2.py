"""
Advent of Code 2024
Day 2: Red-Nosed Reports

Usage:
    python3 day2.py puzzle1 < input
    python3 day2.py puzzle2 < input
"""

import argparse
import sys
from collections.abc import Iterable

Report = tuple[int, ...]


def main() -> None:
    parser = argparse.ArgumentParser()
    parser.add_argument("puzzle", choices=["puzzle1", "puzzle2"])
    args = parser.parse_args()
    if args.puzzle == "puzzle1":
        puzzle = safe_report_count
    else:
        puzzle = safe_report_count_with_problem_dampener
    print(puzzle(parse(sys.stdin.readlines())))


def parse(lines: Iterable[str]) -> Iterable[Report]:
    return (tuple(map(int, line.split())) for line in lines)


def safe_report_count(reports: Iterable[Report]) -> int:
    return sum(1 for report in reports if is_safe(report))


def is_safe(report: Report) -> bool:
    if (
        tuple(sorted(report)) != report
        and tuple(sorted(report, reverse=True)) != report
    ):
        return False
    for i, level in enumerate(report[1:], 1):
        if not 1 <= abs(level - report[i - 1]) <= 3:
            return False
    return True


def safe_report_count_with_problem_dampener(reports: Iterable[Report]) -> int:
    return sum(1 for report in reports if is_safe_with_problem_dampener(report))


def is_safe_with_problem_dampener(report: Report) -> bool:
    if is_safe(report):
        return True
    for i in range(len(report)):
        if is_safe(report[:i] + report[i + 1 :]):
            return True
    return False


if __name__ == "__main__":
    main()
