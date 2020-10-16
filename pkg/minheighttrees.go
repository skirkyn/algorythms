package algorythms

import "math"

//https://leetcode.com/problems/minimum-height-trees/

func findMinHeightTrees(n int, edges [][]int) []int {
	res := make([]int, 0)
	if n == 0 {
		return res
	}
	graph := make([][]bool, n, n)
	for _, edge := range edges {
		row := edge[0]
		col := edge[1]
		if graph[row] == nil {
			graph[row] = make([]bool, n)
		}
		if graph[col] == nil {
			graph[col] = make([]bool, n)
		}
		graph[row][col] = true
		graph[col][row] = true
	}
	min := math.MaxInt32
	minSl := make([]int, 0)

	globalVisited:= make([]bool, n)
	for index := 0; index < n; index++ {
		if !globalVisited[index]{
			visited := make([]bool, n)
			currMin := depth(&visited, &graph, index)
			if currMin < min {
				min = currMin
				minSl = []int{index}
			} else if currMin == min {
				minSl = append(minSl, index)
			}
		}

	}
	return minSl

}

func depth(visited *[]bool, graph *[][]bool, vertex int) int {

	(*visited)[vertex] = true
	minDepth := 0
	for i, v := range (*graph)[vertex] {
		if v && !(*visited)[i] {
			currDepth := depth(visited, graph, i)
			currDepth++
			if currDepth > minDepth {
				minDepth = currDepth
			}
		}
	}
	return minDepth
}
