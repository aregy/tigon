#include<iostream>
#include<vector>
#include<utility>

using std::cin;
using std::cout;
using std::vector;
using std::exception;
using std::swap;

int64_t getMaxProduct(const vector<int>&);

int main(){
  int n;
  cin >> n;
  auto nums = vector<int>(n);
  for (int i = 0; i < n; ++i){
	cin >> nums[i];
  }
  cout << getMaxProduct(nums);
  return 0;
}
int64_t getMaxProduct(const vector<int>& nums){
  auto n = nums.size();
  if (n < 2) {
	throw exception();
  }
  auto start = int(1);
  vector<int> pts(4,0);
  
  for (int el : nums){
	if (start) {
	  start &= 0;
	  pts[1] = el;
	  pts[3] = el;
	}
	
	if (el > pts[1]) {
	  pts[1] = el;
	}
	if (pts[1] > pts[0]) {
	  swap(pts[0], pts[1]);
	}
	
	if (el < pts[3]) {
	  pts[3] = el;
	}
	if (pts[3] < pts[2]){
	  swap(pts[3], pts[2]);
	}
  }
  auto prod1 = (int64_t) pts[0]*pts[1];
  auto prod2 = (int64_t) pts[2]*pts[3];
  return (long long) prod1 < prod2 ? prod2 : prod1;
}
