package algorythms

type FindElements struct {
	root *TreeNode
}

func Constructor(root *TreeNode) FindElements {
	if root == nil {
		return FindElements{}
	}
	if root.Val < 0 {
		root.Val = 0
	}
	if root.Left != nil {
		root.Left.Val = root.Val*2 + 1
	}
	if root.Right != nil {
		root.Right.Val = root.Val*2 + 2
	}
	Constructor(root.Left)
	Constructor(root.Right)
	if root.Val == 0 {
		return FindElements{root: root}
	} else {
		return FindElements{}
	}
}

func (this *FindElements) Find(target int) bool {
	steps := this.buildSteps(target)
	return this.find(len(steps)-1, steps, this.root)
}

func (this *FindElements) find(index int, steps []int, start *TreeNode) bool {
	if index == 0 {
		return start != nil && steps[index] == start.Val
	}
	if start == nil || start.Val != steps[index] {
		return false
	}

	return this.find(index-1, steps, start.Left) || this.find(index-1, steps, start.Right)
}

func (this *FindElements) buildSteps(target int) []int {
	res := make([]int, 0)
	res = append(res, target)
	counter := target
	for counter > 0 {
		var curr int
		if counter%2 == 0 {
			curr = (counter - 2) / 2
		} else {
			curr = (counter - 1) / 2
		}
		counter = curr
		res = append(res, curr)
	}
	return res
}
