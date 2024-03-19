#include <bits/stdc++.h>
// print findKthLargest([3, 5, 2, 4, 6, 8], 3)
void max_heapify(int* nums, size_t len, size_t head){
  if (head*2+1 < len) {
    max_heapify(nums, len, head*2 + 1);
  }
  if (head*2+2 < len) {
    max_heapify(nums, len, head*2 + 2);
  }
  if (head != 0 && nums[(head-1)/2]<nums[head]){
    std::swap(nums[(head-1)/2], nums[head]);
  }
}
int findKthLargest(int* nums, size_t len, int k) {
  for (int i = 0; i < k; i++){
    max_heapify(nums+i, len-i, i);
  }

  return nums[k-1];
}

int main(){
  int n[] = {3, 5, 2, 4, 6, 8};
  std::cout << findKthLargest(n, sizeof(n)/sizeof(n[0]), 3) << std::endl;

  for (const auto& el: n)
    std::cout << el << " ";
  std::cout << std::endl;
  return 0;
}
