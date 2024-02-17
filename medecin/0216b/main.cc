#include <unordered_map>
#include <iostream>
#include <string>

struct Solution
{
    std::unordered_map<char, int> magazine;

public:
    void Magazine()
    {
        std::cout << " ";
        for (auto &[key, value] : this->magazine)
        {
            std::cout << key << ":" << value << ", ";
        }
        std::cout << std::endl;
    }

    bool canSpell(const std::string &mag, const std::string &word)
    {
        magazine.clear();
        for (const char &letter : mag)
        {
            if (magazine.contains(letter))
            {
                magazine[letter]++;
                continue;
            }
            magazine[letter] = 0;
        }
        for (const char &letter : word)
        {
            if (magazine.contains(letter))
            {
                magazine[letter]--;
                if (!magazine[letter])
                    magazine.erase(letter);
                continue;
            }
            return false;
        }
        return true;
    }
};

int main()
{
    std::string mag{"wordee"};
    std::string candidate;
    Solution s;
    while (1)
    {
        std::cout << "Word: ";
        std::cin >> candidate;
        if (s.canSpell(mag, candidate))
            std::cout << "\tCan spell '" << candidate << "'\n";
        else
            std::cout << "\tCANNOT spell '" << candidate << "'\n";
        s.Magazine();
    }
    return 0;
}