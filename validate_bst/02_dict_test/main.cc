#include <iostream>
#include <string>
#include <set>
#include <unordered_map>
#include <memory>

std::pair<bool, std::string> result(const std::string& seed, const std::string& candidate) {
	std::multiset<char> components;
	for (const char& letter : seed)
		components.emplace(letter);
	std::multiset<char>::const_iterator it;
	for (const char& letter : candidate)
		if ((it = components.find(letter)) == components.end())
			return std::make_pair<bool, std::string>(false, "Not a chance");
		else
			components.erase(it);
	return std::make_pair<bool, std::string>(true, "It could happen");
}


bool bagCheck(const std::string& seed, const std::string& candidate) {
	auto mapStore = std::make_unique<std::unordered_map<char, int>>();
	auto& map = *mapStore;
	for (const char& letter : seed)
		if (map.find(letter) == map.end())
			map[letter] = 1;
		else
			map[letter]++;
	std::unordered_map<char, int>::iterator it;
	for (const char& letter : candidate)
		if ((it = map.find(letter)) == map.end())
			return false;
		else
			map[it->first]--;
	return true;
}

int main(){
	std::string seed, candidate;
	std::pair<bool, std::string> response;
	while(1){
		std::cout << "seed: ";
		std::cin >> seed;
		std::cout << "test: ";
		std::cin >> candidate;
		response = result(seed, candidate);
		bool dCheck = bagCheck(seed, candidate);
		if (dCheck != response.first)
			throw candidate + " in " + seed + "???\n";
		if (response.first)
			std::cout << response.second << std::endl;
		else
			std::cout << response.second << std::endl;
	}
}
