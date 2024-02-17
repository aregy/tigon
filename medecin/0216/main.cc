#include <stdio.h>
#include <vector>
#include <iostream>
#include <utility>

// a non-decreasing list is monotonically increasing
template <typename T>
int *indexFromNdcList(const std::vector<T> &nums)
{
    int* pos = new int;
    *pos = -1;
    for (int i = 1; i < nums.size(); i++)
    {
        if (nums[i] < nums[i - 1])
        {
            if (*pos > -1)
            {
                return nullptr;
            }
            *pos = i;
        }
    }
    if (*pos == -1)
        *pos = nums.size();
    return pos;
}
int main()
{
    const std::vector<int> nums{1,
                                2,
                                2,
                                1,
                                3,
                                5,
                                7};
    int *pi = indexFromNdcList(nums);
    if (pi == nullptr){
        std::cout << "Non-modifiable by 1 to reach condition\n";
        return 0;
    }
    if (*pi == nums.size()){
        std::cout << "List is already monotonically increasing";
        return 0;
    }
    std::cout << printf("Modify %d (@%d)\n", nums[*pi], *pi);
    std::cout << "\t";
    for (int i = 0; i < nums.size(); i++){
        if (i == *pi){
            std::cout << nums[i] << "! ";
            continue;
        }
        std::cout << nums[i] << " ";
    }
    std::cout << std::endl;

    return 0;
}