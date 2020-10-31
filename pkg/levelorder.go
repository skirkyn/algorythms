package algorythms

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	traverse(root, &res, 0)
	return res
}

func traverse(root *TreeNode, arr *[][]int, level int) {
	if root != nil {
		if len(*arr) <= level {
			*arr = append(*arr, []int{})
		}
		(*arr)[level] = append((*arr)[level], root.Val)
		traverse(root.Left, arr, level+1)
		traverse(root.Right, arr, level+1)
	}
}
