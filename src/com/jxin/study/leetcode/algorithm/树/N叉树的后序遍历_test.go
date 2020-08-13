package 树

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
