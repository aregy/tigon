#pragma once

#include <cxxabi.h>
#include <iostream>
#include <string>
#include <type_traits>
#include <typeinfo>
#include <utility>

#define _ArY_type(expr)                                                        \
  do {                                                                         \
    using Type = std::remove_reference<decltype(expr)>;                        \
    bool isVolatile = std::is_volatile<Type>::value;                           \
    bool isConst = std::is_const<Type>::value;                                 \
    bool isArraylike = std::is_array<Type>::value;                             \
    bool isRValueReference = std::is_rvalue_reference<decltype(expr)>::value;  \
    bool isLValueReference = std::is_lvalue_reference<decltype(expr)>::value;  \
    std::cout << expr << " (";                                                 \
    if (isVolatile)                                                            \
      std::cout << "volatile ";                                                \
    if (isConst)                                                               \
      std::cout << "const ";                                                   \
    int statusCode;                                                            \
    char *demangle =                                                           \
        abi::__cxa_demangle(typeid(decltype(expr)).name(), 0, 0, &statusCode); \
    std::cout << demangle;                                                     \
    if (isRValueReference)                                                     \
      std::cout << "&&";                                                       \
    if (isLValueReference)                                                     \
      std::cout << "&";                                                        \
    if (isArraylike)                                                           \
      std::cout << "[]";                                                       \
    std::cout << ")\n";                                                        \
    free(demangle);                                                            \
  } while (0)
namespace ArY {

template <typename T> struct repr {
  static std::string to_string();
};
} // namespace ArY
template <typename T> std::string ArY::repr<T>::to_string() {
  int statusCode;
  std::unique_ptr<char, void (*)(void *)> repr{
      abi::__cxa_demangle(typeid(T).name(), 0, 0, &statusCode), std::free};
  if (statusCode)
    return typeid(T).name();
  return repr.get();
}
