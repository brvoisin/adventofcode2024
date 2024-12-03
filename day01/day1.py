"""
Advent of Code 2024
Day 1: Historian Hysteria

Usage:
    python3 day1.py puzzle1 < input
    python3 day1.py puzzle2 < input
"""

import argparse
import sys
from collections import Counter
from collections.abc import Generator
from collections.abc import Iterable


def main() -> None:
    parser = argparse.ArgumentParser()
    parser.add_argument("puzzle", choices=["puzzle1", "puzzle2"])
    args = parser.parse_args()
    if args.puzzle == "puzzle1":
        puzzle = total_distance
    else:
        puzzle = total_similarity_score
    print(puzzle(parse(sys.stdin.readlines())))


def parse(lines: Iterable[str]) -> Iterable[tuple[int, int]]:
    return (parse_pair(line.strip()) for line in lines)


def parse_pair(pair: str) -> tuple[int, int]:
    return tuple(map(int, pair.split()))


def total_distance(pair_list: Iterable[tuple[int, int]]) -> int:
    """Puzzle 1"""
    return sum(distances(pair_list))


def distances(pair_list: Iterable[tuple[int, int]]) -> Generator[int, None, None]:
    left, right = unzip(pair_list)
    for location_1, location_2 in zip(sorted(left), sorted(right)):
        yield abs(location_1 - location_2)


def unzip(pair_list: Iterable[tuple[int, int]]) -> tuple[Iterable[int], Iterable[int]]:
    return zip(*pair_list, strict=True)


def total_similarity_score(pair_list: Iterable[tuple[int, int]]) -> int:
    """Puzzle 2"""
    return sum(similarity_scores(pair_list))


def similarity_scores(
    pair_list: Iterable[tuple[int, int]]
) -> Generator[int, None, None]:
    left, right = unzip(pair_list)
    right_counter = Counter(right)
    for location in left:
        yield location * right_counter[location]


if __name__ == "__main__":
    main()
