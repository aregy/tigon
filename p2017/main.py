class Solution:
    def gridGame(self, grid: list[int]) -> list[int]:
        probSpace = Solution.takePaths(grid)
        slnSpace = [0 for k in probSpace]
        print('probSpace:', probSpace)
        for idx, prob in enumerate(probSpace):
            slnSpace[idx] = Solution.bestPath(prob)

        return slnSpace

    @classmethod
    def takePaths(cls, grid: list[list[int]]) -> list[list[int]]:
        slns = []
        for i in range(len(grid[0])):
            path = []
            for idx, _ in enumerate(grid):
                path.append(grid[idx][:])

            for k, _ in enumerate(grid[0]):
                if k > i:
                    break
                path[0][k] = 0
            for k, _ in enumerate(grid[1]):
                if k < i:
                    continue
                path[1][k] = 0

            slns.append(path)
        return slns

    @staticmethod
    def bestPath(grid: list[list[int]]) -> int:
        routeSums = [0] * len(grid[0])
        for idx, _ in enumerate(routeSums):
            pos = (0, 0)
            tick = 0
            while pos[0] < len(grid) - 1 and pos[1] < len(grid[0]) - 1:
                routeSums[idx] += grid[pos[0]][pos[1]]
                if tick == idx:
                    pos = (pos[0] + 1, pos[1])
                else:
                    pos = (pos[0], pos[1] + 1)
                tick += 1
        return max(routeSums)


if __name__ == '__main__':
    testcases = list()
    testcases.append([2,5,4], [1,5,1])
    # testcases.append([[1, 2], [3, 4]])
    # testcases.append([[2, 5, 4], [1, 5, 1]])
    # testcases.append([[3, 3, 1], [8, 5, 2]])
    # testcases.append([[1, 3, 1, 15], [1, 3, 3, 1]])
    for idx, test in enumerate(testcases):
        print(f"Testcase #{idx + 1}: given {test}...")
        print(Solution().gridGame(test))

