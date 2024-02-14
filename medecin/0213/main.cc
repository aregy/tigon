// in-order traversal of a tree
#include <vector>
#include <iostream>

template <typename T>
struct Node
{
    T val;
    Node<T> *left;
    Node<T> *right;
    Node(T v, Node<T> *l = nullptr, Node<T> *r = nullptr) : val(v), left(l), right(r) {}
};

template <typename T>
void getDepth(const Node<T> &node, int &depth)
{
    if (node.right || node.left)
        depth += 1;
    if (node.right)
        getDepth<int>(*node.right, depth);
    if (node.left)
        getDepth<int>(*node.left, depth);
}
template <class T>
void PrintTree(Node<T> &root)
{
    auto cRow = std::vector<Node<T>>{root};
    std::vector<Node<T>> nRow{};
    int level = 0;
    int depth = 0;
    getDepth<int>(root, depth);
    while (cRow.size() != 0)
    {
        while (cRow.size() != 0)
        {
            for (int i = 0; i < depth - level; i++)
                std::cout << " ";
            auto node = cRow.back();
            std::cout << node.val;
            cRow.pop_back();
            if (node.left)
                nRow.insert(nRow.begin(), *node.left);
            if (node.right)
                nRow.insert(nRow.begin(), *node.right);
        }
        std::cout << "\n";
        level++;
        cRow = nRow;
        nRow = {};
    }
}
template <typename T>
void InOrderTraversal(Node<T> root, void (*action)(const Node<T>&))
{
    std::vector <std::pair<bool, Node<T>>> path{};
    path.push_back(std::make_pair(false, root));
    while (path.size())
    {
        auto [expanded, node] = path.back();
        path.pop_back();
        if (expanded) {
            action(node);
            continue;
        }
        if (node.right){
            path.push_back(std::make_pair(false, *node.right));
        }
        path.push_back(std::make_pair(true, node));
        if (node.left){
            path.push_back(std::make_pair(false, *node.left));
        }
    }
}
template<typename T>
void printNode(const Node<T>& node){
    std::cout << node.val << " ";
}
int main()
{
    Node<int> n = Node<int>(12, new Node<int>(6, new Node<int>(2), new Node<int>(3)), new Node<int>(4, new Node<int>(7), new Node<int>(8)));
    PrintTree(n);
    std::cout << "\n";
    InOrderTraversal(n, printNode);
    std::cout << "\n";
    
    return 0;
}