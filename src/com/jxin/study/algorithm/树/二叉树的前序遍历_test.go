package 树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret []int

func preorderTraversal(root *TreeNode) []int {
	ret = []int{}
	frontFirst(root)
	return ret
}
func frontFirst(node *TreeNode) {
	if node == nil {
		return
	}
	ret = append(ret, node.Val)
	frontFirst(node.Left)
	frontFirst(node.Right)
}
