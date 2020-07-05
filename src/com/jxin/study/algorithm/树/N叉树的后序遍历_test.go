package 树

type Node struct {
	Val      int
	Children []*Node
}

var ret []int

func postorder(root *Node) []int {
	ret = []int{}
	addNode(root)
	return ret
}

func addNode(node *Node) {
	if node == nil {
		return
	}

	for _, v := range node.Children {
		addNode(v)
	}
	ret = append(ret, node.Val)
}
