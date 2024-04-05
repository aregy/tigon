		#include "pch/bstdc++.h"

		std::unordered_map<unsigned int, unsigned int> _cache = {{1, 1}, {2, 2}};

		unsigned int cached_staircase(unsigned int n)
		{
			if (_cache.find(n) != _cache.end())
				return _cache[n];
			if (n - 2 > 0)
				return cached_staircase(n - 1) + cached_staircase(n - 2);
			else if (n - 1 > 0)
				return cached_staircase(n - 1);
			return 1;
		}

		unsigned int staircase(unsigned int n)
		{
			std::unordered_multiset<unsigned int> positions = {n};
			bool processed;
			#if DEBUG
			auto print = [&positions]
			{
				for (const auto el : positions)
				{
					std::cout << el << " ";
				}
				std::cout << "\n";
			};
			int i = 0;
			#endif
			while (!processed)
			{
				processed = true;
			#if DEBUG
				std::cout << "Iteration #" << ++i << ": ";
				print();
				std::cout << '\n';
			#endif
				for (auto el : positions)
				{
					if (el == 0)
						continue;
					processed = false;
					positions.erase(positions.find(el));
					positions.insert(el - 1);
					if (el > 1)
						positions.insert(el - 2);
				}
			}
			return positions.size();
		}

		int main(int argc, char* argv[])
		{
			auto n = stoull(static_cast<std::string>(argv[1]));
			std::cout << n << "-> "<< staircase(n) << '\n';
			return 0;
		}
