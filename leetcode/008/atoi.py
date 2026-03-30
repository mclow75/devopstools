import unittest


class Solution:
    def my_atoi(self, s: str) -> int:
        s = s.strip()
        if not s:
            return 0
        sign, result = 1, 0
        if s[0] in "-+":
            if s[0] == "-":
                sign = -1
            s = s[1:]
        for char in s:
            if not char.isdigit():
                break
            result = result * 10 + int(char)

        result *= sign
        if result < -(2**31):
            return -(2**31)
        elif result > 2**31 - 1:
            return 2**31 - 1
        return result


class TestSolution(unittest.TestCase):
    def test_add(self):
        self.assertEqual(add(1, 2), 3)


if __name__ == "__main__":
    unittest.main()
