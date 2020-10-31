package algorythms

import (
	"math"
	"strconv"
)

func printTree(root *TreeNode) [][]string {
	theMap := make(map[int]map[int]int, 0)
	width, depth := traversePrint(root, 0, 0, &theMap)
	res := make([][]string, depth, width)
	for i := 0; i < depth; i++ {
		for j := 0; j < width; j++ {
			if row, ok := theMap[i]; ok {
				if val, ok := row[j]; ok {
					res[i][j] = strconv.Itoa(val)
				} else {
					res[i][j] = ""
				}
			} else {
				res[i][j] = ""
			}
		}
	}
	return res
}

func traversePrint(root *TreeNode, shift, depth int, theMap *map[int]map[int]int) (int, int) {
	if root != nil {
		if val, ok := (*theMap)[depth]; ok {
			val[shift] = root.Val
			(*theMap)[depth] = val
		} else {
			(*theMap)[depth] = make(map[int]int, 0)
			(*theMap)[depth][shift] = root.Val
		}
		leftShift, leftDepth := traversePrint(root.Left, shift+1, depth+1, theMap)
		rightShift, rightDepth := traversePrint(root.Right, shift+1, depth+1, theMap)
		return 1 + leftShift + rightShift, int(math.Max(float64(leftDepth), float64(rightDepth)))
	}
	return 0, depth
}
