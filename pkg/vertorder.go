package algorythms

import (
	"math"
)

func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	mapping := make(map[int][]int, 0)
	queue := make(map[int][]*TreeNode, 0)
	queue[0] = []*TreeNode{root}
	min := math.MaxInt32
	max := math.MinInt32
	next := []int{0}
	for len(next) > 0 {
		nextNum := next[0]
		min = int(math.Min(float64(min), float64(nextNum)))
		max = int(math.Max(float64(max), float64(nextNum)))
		next = next[1:]
		nodes := queue[nextNum]
		delete(queue, nextNum)
		for _, v := range nodes {
			if v.Left != nil {
				leftNum := nextNum - 1
				next = append(next, leftNum)
				if arr, ok := queue[leftNum]; ok {
					arr = append(arr, v.Left)
					queue[leftNum] = arr
				} else {
					queue[leftNum] = []*TreeNode{v.Left}
				}
			}
			if v.Right != nil {
				rightNum := nextNum + 1
				next = append(next, rightNum)
				if arr, ok := queue[rightNum]; ok {
					arr = append(arr, v.Right)
					queue[rightNum] = arr
				} else {
					queue[rightNum] = []*TreeNode{v.Right}
				}
			}
			if arr, ok := mapping[nextNum]; ok {
				mapping[nextNum] = append(arr, v.Val)
			} else {
				mapping[nextNum] = []int{v.Val}
			}
		}

	}
	length := int(math.Abs(float64(min)) + math.Abs(float64(max)) + 1)
	arr := make([][]int, length)
	for k, v := range mapping {
		index := int(math.Abs(float64(min - k)))
		arr[index] = v
	}
	return arr
}

