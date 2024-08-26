#include <iostream>
#include <string>
#include <concepts>
#include <limits>
#include <algorithm>

template <typename T>
concept Numeric = requires(T a) {
	{ std::numeric_limits<T>::min() } -> std::convertible_to<T>;
	{ std::numeric_limits<T>::max() } -> std::convertible_to<T>;
	{ std::cout << a } -> std::same_as<std::ostream&>;
};
template <Numeric T>
struct Node {
	T val;
	Node<T>* left = nullptr;
	Node<T>* right = nullptr;
	Node() = default;
	Node(T t1) : val(t1) {}
	void Print() const {
		std::cout << repr() << std::endl;
	}
	std::string repr() const {
		std::string s = "";
		s += "(";
		s += std::to_string(val);
		if (left != nullptr || right != nullptr)
		 	s += " ";
		if (left != nullptr){
			s += left->repr();
			if (right != nullptr)
				s += " ";
		}
		if (right != nullptr)
			s += right->repr();
		s += ")";
		return s;
	}
};
template <Numeric T>
bool IsValidBST(Node<T>* n, T floor = std::numeric_limits<T>::min(), T ceil = std::numeric_limits<T>::max()){
	if (n == nullptr)
		return true;
	if (n->val < floor || n->val > ceil)
		return false;

	T newFloor = floor, newCeil = ceil;
	if (n->val > newFloor)
			newFloor = n->val;
	if (n->val < newCeil)
			newCeil = n->val;
	
	return IsValidBST(n->left, floor, newCeil) && IsValidBST(n->right, newFloor, ceil);
}
typedef Node<int> iNode;
int main(){
	iNode n(10), n2(4), n3(11), n4(7);
	n.left = &n2;
	n.right = &n3;
	n.Print();
	std::cout << IsValidBST(&n) << std::endl;
	n3.left = &n4;
	n.Print();
	std::cout << IsValidBST(&n) << std::endl;
	std::cout << std::endl;
	return 0;
}
