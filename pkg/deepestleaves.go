package algorythms

//https://leetcode.com/problems/deepest-leaves-sum/

func deepestLeavesSum(root *TreeNode) int {

	sums := make([]int, 0)
	deepest := sum(root, 0, &sums)
	if deepest > 0 {
		return sums[deepest]
	} else {
		return 0
	}
}

func sum(root *TreeNode, depth int, sums *[]int) int {

	if root != nil {
		if len(*sums) <= depth {
			*sums = append(*sums, 0)
		}
		if root.Right == nil && root.Left == nil {
			(*sums)[depth] += root.Val
			return depth
		} else {
			left := sum(root.Left, depth+1, sums)
			right := sum(root.Right, depth+1, sums)
			if left > right {
				return left
			} else {
				return right
			}
		}
	}
	return -1
}
