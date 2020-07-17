package 树

func inorderTraversal(root *TreeNode) []int {
	ret = []int{}
	middleFirst(root)
	return ret
}

func middleFirst(node *TreeNode) {
	if node == nil {
		return
	}
	middleFirst(node.Left)
	ret = append(ret, node.Val)
	middleFirst(node.Right)
}
