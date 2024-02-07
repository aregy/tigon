#include <iostream>
#include <string>
#include <utility>

static std::string str1;
void revString(const std::string::iterator begin, const std::string::iterator &end)
{
    auto front = begin;
    auto back = end;
    --back;
    int i = 0;
    long int ri;
    while (true)
    {
        std::swap(*front, *back);
        i++;
        ri = str1.size() - 1 - i;
        ++front;
        if (front == back)
            break;
        --back;
        if (front == back)
            break;
    }
}
int main()
{
    str1 = std::string{"This should get reversed. And it does."};
    std::cout << str1 << std::endl;
    revString(str1.begin(), str1.end());
    std::cout << str1 << std::endl;
}
