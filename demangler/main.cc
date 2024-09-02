#include "defines.hpp"

int main() {
  int x = 10;
  const int y = 20;
  int arr[5] = {1, 2, 3, 4, 5};
  int &rf = x;
  _T(x);
  _T(std::move(x));
  _T(y);
  _T(arr);
  _T(42);
  _T(rf);

  return 0;
}
