from typing import List


class Solution:
    def two_sum(self, nums: List[int], target: int) -> List[int]:
        known_numbers = {}
        for idx, num in enumerate(nums):
            desire = target - num
            if desire in known_numbers:
                return [known_numbers[desire], idx]
            known_numbers[num] = idx
        return [0, 0]


if __name__ == "__main__":
    s = Solution()
    assert s.two_sum([2, 7, 11, 15], 9) == [0, 1], "Результат неверен"
    # assert s.two_sum([3, 2, 4], 6)
    # assert add(2, 3) == 5, "The function did not return the expected value"
