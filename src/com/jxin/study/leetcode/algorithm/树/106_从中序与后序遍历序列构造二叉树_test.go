package 树

import "testing"

func TestService2(t *testing.T) {

	buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})

}
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{postorder[len(postorder)-1], nil, nil}
	i := 0
	for i < len(inorder) && postorder[len(postorder)-1] != inorder[i] {
		i++
	}

	root.Left = buildTree(inorder[:i], postorder[:len(postorder)-1-len(inorder[i+1:])])
	root.Right = buildTree(inorder[i+1:], postorder[len(postorder)-1-len(inorder[i+1:]):len(postorder)-1])
	return root
}
