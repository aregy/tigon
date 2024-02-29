#include <iostream>
using std::cout;

template<typename T>
class Wrapper {
    T m_Value;
public:
    Wrapper(T value) : m_Value(value) { }
    Wrapper operator++(int) {
        Wrapper<T> previous{ this->m_Value };
        this->m_Value = this->m_Value + 1;
        return previous;
    }
    //std::ostream& operator<<(std::ostream&, Wrapper<T>&);

    std::ostream& operator<<(std::ostream& o);
    T& GetValue();
    bool operator< (Wrapper b) const;
};

template<typename T>
T& Wrapper<T>::GetValue() {
    return m_Value;
}
template <typename T>
std::ostream& Wrapper<T>::operator<<(std::ostream& o) {
    T& tVal = this->GetValue();
    o << tVal;
    return o;
}
template <typename T>
std::ostream& operator<<(std::ostream& o, Wrapper<T>& wrapper){
    return wrapper.operator<<(o);
}

// without the following template we get...
// main.cc:(.text+0x84): undefined reference to `Wrapper<int>::operator<(Wrapper<int>) const'
// collect2: error: ld returned 1 exit status

template<typename T>
bool Wrapper<T>::operator< (Wrapper b) const {
    return m_Value < b.GetValue();
}

int main() {
    for (Wrapper<int> i = 0; i < 5; i++) {
        std::cout << i << " ";
    }
    return 0;

}

