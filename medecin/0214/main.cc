#include <iostream>
#include <vector>
#include <string>
#include <algorithm>

template<typename T>
void partition(std::vector<T>& items, int l, int r){
	if (r <= l + 1)
		return;
	int p = l;
	for (int i = l+1; i < r; i++){
		if (items[i] < items[p]){
			std::swap(items[i], items[p]);
			p++;
		}	
	}
	partition(items, l, p);
	partition(items, p+1, r);
}	
void partition(std::string& str, int l, int r){
	if (r <= l + 1)
		return;
	int p = l;
	for (int i = l+1; i < r; i++){
		if (str[i] < str[p]){
			std::swap(str[i], str[p]);
			p++;
		}	
	}
	partition(str, l, p);
	partition(str, p+1, r);
}	
void qsort(std::string& str){
	partition(str, 0, str.length());
}
void sayBack(std::string str){
	for (auto it = str.begin(); it != str.end(); ++it){
		std::cout << *it;
	}
	std::cout << std::endl;
}
std::string derangement(std::string& str){
	qsort(str);
	int f = 2;
	while (f){
	for (int i = 1; i < str.length(); i++){
		if (str[i-1] == str[i])
			std::swap(str[i], str[(i+1)%str.length()]);
	}
	f--;
	}
	for (int i = 1; i < str.length(); i++){
		if (str[i-1] == str[i]) {
			return "NO, EVEN " + str + " FAILS\n";
		}
	}
	return str;
}
int main(){
	//std::vector nums {10, 9, 8, 7, 6, 5, 4, 3, 2, 1};
	//partition(nums, 0, nums.size());
	//for (auto el : nums)
		//std::cout << el << " ";
	//std::cout << std::endl;
	//return 0;
	std::string word;
	std::cout << ": ";
	std::cin >> word;
	std::cout << derangement(word);
	return 0;
}
