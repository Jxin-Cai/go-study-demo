package 树

import "testing"

/*type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}*/
func TestService(t *testing.T) {

	buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})

}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for i < len(inorder) && preorder[0] != inorder[i] {
		i++
	}

	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
