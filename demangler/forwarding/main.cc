#include "../defines.hpp"
#include <iostream>
#include <utility>

// template <typename T> void p(T t) { std::cout << "Forwarded as T value\n"; }
template <typename T> void p(T &t) { std::cout << "Forwarded as T& value\n"; }
template <typename T> void p(T &&t) { std::cout << "Forwarded as T&& value\n"; }

template <typename T> void u(T &&t) {
  _ArY_type(t);
  std::cout << ArY::repr<T>::to_string() << std::endl;
}
class C {
  int X{};
};
int main() {
  u(std::forward<int>(100)); // forwarded as int&& (I think)
  u(100);                    // forwarded as int&& (I think)
  C c1;
  C &c2 = c1;
  std::cout << ArY::repr<decltype(c2)>::to_string() << std::endl;
  return 0;
}
