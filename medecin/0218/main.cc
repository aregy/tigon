#include <algorithm>
#include <iostream>
#include <utility>
#include <vector>

template <typename T> void rotate_list(std::vector<T> &list, int k) {
  k %= list.size();
  if (k == 0)
    return;
  std::reverse(list.begin(), list.end());
  std::reverse(list.begin(), list.begin() + k);
  std::reverse(list.begin() + k, list.end());
}

int main() {
  auto list =
      std::vector<char>{'1', '2', '3', '4', '5', '6', '7', 'a', 'b', 'c'};
  for (const char &el : list)
    std::cout << el << " ";
  std::cout << std::endl;
  rotate_list<char>(list, 3);
  for (const char &el : list)
    std::cout << el << " ";
  std::cout << std::endl;
}