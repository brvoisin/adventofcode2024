from day3 import sum_multiplicaton_results

sample = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"


def test_sum_multiplication_results():
    assert sum_multiplicaton_results(sample) == 161
