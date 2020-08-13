package 树

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
