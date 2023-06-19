#include<iostream>


int main(){
  int nums[] = {0};
  int len = sizeof(nums)/sizeof(nums[0]);
  for (int i = 0; i < len; ++i){
    nums[i] = len - 1 - i;
  }
  for (int i = 0; i < len; ++i){
    std::cout << nums[i] << ", ";
  }
  std::cout << std::endl;
}
