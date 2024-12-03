from day1 import parse, unzip, total_distance, total_similarity_score

sample = """3   4
4   3
2   5
1   3
3   9
3   3"""

pairs = [
    (3, 4),
    (4, 3),
    (2, 5),
    (1, 3),
    (3, 9),
    (3, 3),
]


def test_parse():
    assert list(parse(sample.split("\n"))) == pairs


def test_unzip():
    left, right = unzip(pairs)
    assert list(left) == [3, 4, 2, 1, 3, 3]
    assert list(right) == [4, 3, 5, 3, 9, 3]


def test_total_distance():
    assert total_distance(pairs) == 11


def test_similarity_score():
    assert total_similarity_score(pairs) == 31
