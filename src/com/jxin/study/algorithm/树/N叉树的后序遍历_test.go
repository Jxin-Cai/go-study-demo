package 树

type Node struct {
	Val      int
	Children []*Node
}

var ret []int

func postorder(root *Node) []int {
	ret = []int{}
	laterNode(root)
	return ret
}

func laterNode(node *Node) {
	if node == nil {
		return
	}

	for _, v := range node.Children {
		laterNode(v)
	}
	ret = append(ret, node.Val)
}
