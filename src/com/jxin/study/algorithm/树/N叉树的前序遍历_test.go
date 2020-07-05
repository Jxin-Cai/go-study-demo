package 树

type Node struct {
	Val      int
	Children []*Node
}

var ret []int

func preorder(root *Node) []int {
	ret = []int{}
	frontNode(root)
	return ret
}

func frontNode(node *Node) {
	if node == nil {
		return
	}
	ret = append(ret, node.Val)
	for _, v := range node.Children {
		frontNode(v)
	}
}
