package 树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret []int

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
