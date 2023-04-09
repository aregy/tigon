#!/usr/bin/env python3
def maxComp(a, b):
    if a > b:
        return -1
    elif a == b:
        return 0
    return 1

def minComp(a, b):
    if a < b:
        return -1
    elif a == b:
        return 0
    return 1

def hfyDown(nums, i, comp=minComp):
    c = i

    if comp(nums[i], nums[2*i + 1]) == 1:
        c = 2*i + 1
    if comp(nums[c], nums[2*i + 2]) == 1:
        c = 2*i + 2

    if i != c:
        nums[c], nums[i] = nums[i], nums[c]

    if i != 0:
        return 2 + hfyDown(nums, (i - 1)//2, comp)
    return 2

def get_mm(nums):
    if len(nums) == 0:
        return None
    comparisons = 0
    for i in range(len(nums) // 2 - 1, -1, -1):
        comparisons += hfyDown(nums, i)
    min = nums[0]
    max = nums[len(nums) // 2 + 1]
    for i in range(len(nums) // 2 + 1, len(nums)):
        comparisons += 1
        if nums[i] > max:
            min = nums[i]
    return (min, max, comparisons)

def get_min_max(nums):
    n = len(nums)
    if n == 0:
        return None

    if n % 2 == 0:
        min = nums[-1]
        max = nums[-1]
    else:
        min = nums[0]
        max = nums[0]

    comparisons = 0
    for i in range(n // 2):
        comparisons += 3
        if nums[2*i] > nums[2*i + 1]:
            localMax = nums[2*i]
            localMin = nums[2*i + 1]
        else:
            localMin = nums[2*i]
            localMax = nums[2*i + 1]
        if localMax > max:
            max = localMax
        if localMin < min:
            min = localMin

    return (nums, (min, max, comparisons), get_mm(nums))

if __name__ == '__main__':
    print(get_min_max([1, 9, 4, -4, 8, 43, 2, -6, 9]))
