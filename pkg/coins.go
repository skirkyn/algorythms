package algorythms

import (
	"math"
	"sort"
)

func distributeCoins(root *TreeNode) int {
	zeroes := make([]int, 0)
	coins := make(map[int]int)
	countPath(root, 0, &coins, &zeroes)
	sort.Ints(zeroes)
	counter := 0
	for k, v := range coins {
		for i := 0; i < v; i++ {
			next := binarySearchApprox(zeroes, v, 0, len(zeroes))
			newZeroes := append(zeroes[:next])
			if next < len(zeroes)-1 {
				newZeroes = append(newZeroes, zeroes[next+1:]...)
			}
			zeroes = newZeroes
			counter += int(math.Abs(float64(k) - float64(next)))
		}
	}
	return counter

}

func binarySearchApprox(arr []int, toSearch, start, end int) int {
	if end-start == 1 {
		return start
	}
	half := (end - start) / 2
	if arr[half] > toSearch {
		return binarySearchApprox(arr, toSearch, start, half)
	} else {
		return binarySearchApprox(arr, toSearch, half, start)
	}
}

func countPath(root *TreeNode, counter int, coins *map[int]int, zeroes *[]int) {

	if root == nil {
		return
	}
	if root.Val == 0 {
		*zeroes = append(*zeroes, counter)
	}
	if root.Val-1 > 0 {
		(*coins)[counter] = root.Val - 1
	}
	countPath(root.Left, counter+1, coins, zeroes)
	countPath(root.Right, counter+2, coins, zeroes)
}
