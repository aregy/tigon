#include <cstdio>
#include <cstdio>
#include <cmath>
#include <string>
#include <vector>
#include <functional>

int numTrailing(const int& x) {
	int trail = 0;
	int divisor = 10;
	while (x % divisor == 0){
		divisor *= 10;
		trail++;
	}
	return trail;
}
int trailingZeroesOfFactorial(const int& x) {
  std::vector<int> nums(x);
  for (int i = 0; i < x; i++){
	nums[i] = x - i;
  }
  std::function<int(int&, int)> kth_power = [] (int& store, int n) -> int {
	int k_n = 0;
	while (store % n == 0){
	  store /= n;
	  k_n++;
	}
	return k_n;
  };
  int k2 = 0;
  int k5 = 0;
  for (int& el : nums){
	k2 += kth_power(el, 2);
	k5 += kth_power(el, 5);
  }
  int T = static_cast<int>(pow(2, k2) * pow(5, k5));
  return kth_power(T, 10);
}
void printTrailing(const std::string& qty, int trailing) {
  switch (trailing){
	case 0:
	  printf("No trailing 0s for %s\n", qty.c_str());
	  break;
	default:
	  printf("%d trailing 0s for %s\n", trailing, qty.c_str());
  }
}
int main(){
	int x;
	while (1){
		printf(": ");
		scanf("%d", &x);
		printTrailing(std::to_string(x), numTrailing(x));
		printTrailing(std::to_string(x) + "!", trailingZeroesOfFactorial(x));
	}
}
