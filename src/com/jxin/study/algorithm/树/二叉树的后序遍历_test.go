package 树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret []int

func postorderTraversal(root *TreeNode) []int {
	ret = []int{}
	laterFirst(root)
	return ret
}

func laterFirst(node *TreeNode) {
	if node == nil {
		return
	}
	laterFirst(node.Left)
	laterFirst(node.Right)
	ret = append(ret, node.Val)
}
