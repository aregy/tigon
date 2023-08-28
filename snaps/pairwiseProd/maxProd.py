def maxProduct(nums : list[int]) -> int:
    max, max2, min, min2 = 0, 0, 0, 0
    for i, el in enumerate(nums):
        if i == 0:
            max1 = el
            min1 = el
        if el < min:
            min2 = min
            min = el
        elif el < min2:
            min2 = el
            
        if el > max:
            max2 = max
            max = el
        elif el > max2:
            max2 = el

    prod1, prod2 = max * max2, min * min2
    if prod2 > prod1:
        return prod2
    return prod1

_ = input()
nums = list(map(int, input().split()))
print(maxProduct(nums))
