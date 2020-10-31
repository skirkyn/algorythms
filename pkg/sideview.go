package algorythms

func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)
	next := root
	for next != nil {
      result=append(result, next.Val)
      if next.Right != nil{
		  next = root.Right
	  }else {
	  	next = root.Left
	  }
	}
	return result
}
