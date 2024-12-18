import pytest

from day2 import is_safe
from day2 import is_safe_with_problem_dampener
from day2 import parse
from day2 import safe_report_count
from day2 import safe_report_count_with_problem_dampener


sample = """7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9"""

reports = [
    (7, 6, 4, 2, 1),
    (1, 2, 7, 8, 9),
    (9, 7, 6, 2, 1),
    (1, 3, 2, 4, 5),
    (8, 6, 4, 4, 1),
    (1, 3, 6, 7, 9),
]


def test_parse():
    assert list(parse(sample.split("\n"))) == reports


@pytest.mark.parametrize(
    "report,safe",
    [
        ((7, 6, 4, 2, 1), True),
        ((1, 2, 7, 8, 9), False),
        ((9, 7, 6, 2, 1), False),
        ((1, 3, 2, 4, 5), False),
        ((8, 6, 4, 4, 1), False),
        ((1, 3, 6, 7, 9), True),
    ],
)
def test_is_safe(report, safe):
    assert is_safe(report) is safe


def test_safe_report_count():
    assert safe_report_count(reports) == 2


@pytest.mark.parametrize(
    "report,safe",
    [
        ((7, 6, 4, 2, 1), True),
        ((1, 2, 7, 8, 9), False),
        ((9, 7, 6, 2, 1), False),
        ((1, 3, 2, 4, 5), True),
        ((8, 6, 4, 4, 1), True),
        ((1, 3, 6, 7, 9), True),
        ((1, 2, 3, 2, 1), False),
    ],
)
def test_is_safe_with_problem_dampener(report, safe):
    assert is_safe_with_problem_dampener(report) is safe


def test_safe_report_count_with_problem_dampener():
    assert safe_report_count_with_problem_dampener(reports) == 4
