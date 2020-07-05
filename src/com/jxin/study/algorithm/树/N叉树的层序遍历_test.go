package 树

type Node struct {
	Val      int
	Children []*Node
}

var ret [][]int

func levelOrder1(root *Node) [][]int {
	if root == nil {
		return nil
	}

	queue := make([]*Node, 0)
	result := make([][]int, 0)

	queue = append(queue, root)
	for len(queue) != 0 {
		len := len(queue)
		layer := make([]int, 0)
		for i := 0; i < len; i++ {
			node := queue[0]
			queue = queue[1:]

			for _, child := range node.Children {
				queue = append(queue, child)
			}

			layer = append(layer, node.Val)
		}
		result = append(result, layer)
	}
	return result
}

func levelOrder2(root *Node) [][]int {
	if root == nil {
		return nil
	}
	ret = [][]int{}
	levelNode(root, 0)
	return ret
}

func levelNode(node *Node, level int) {
	if node == nil {
		return
	}
	if len(ret) == level {
		ret = append(ret, []int{})
	}
	ret[level] = append(ret[level], node.Val)
	for _, v := range node.Children {
		levelNode(v, level+1)
	}
}
