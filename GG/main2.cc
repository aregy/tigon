#include <algorithm>
#include <iostream>
#include <vector>

class C
{
public:
  C(int s)
  {
    this->size = static_cast<size_t>(s);
    content = new int[s];
  }
  C(const C &c0)
  {
    size = c0.size;
    content = new int[size];
    std::copy(c0.content, c0.content + size, content);
  }
  C(C &&other)
  {
    std::cout << "move constructor called\n";
    size = other.size;
    content = other.content;
    other.size = 0;
    other.content = nullptr;
  }

  C &operator=(const C &right)
  {
    if (this == &right)
      return *this;
    delete[] content;
    size = right.size;
    content = new int[size];
    std::copy(right.content, right.content + size, content);
    return *this;
  }

  C &operator=(C &&right)
  {
    std::cout << "move assignment operator called\n";
    if (this == &right)
      return *this;
    size = right.size;
    delete[] content;
    content = right.content;
    right.size = 0;
    right.content = nullptr;
    return *this;
  }

  ~C()
  {
    delete[] content;
  }

private:
  size_t size;
  int *content;
};
C rvalue(int size)
{
  return C(size);
}
template <typename T>
void qs(std::vector<T> &nums)
{
}
template <class T>
void print(T t)
{
  for (typename T::iterator it = t.begin(); it != t.end(); it++)
  {
    std::cout << *it << ", ";
  }
  std::cout << std::endl;
}
int main()
{
  std::vector<int> n1 = {1, 2, -4, 0, 3, 2, 1, 9, 11, -9, -3, -4};
  qs(n1);
  print(n1);
}