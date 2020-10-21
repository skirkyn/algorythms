package algorythms

import "math"

//https://leetcode.com/problems/path-in-zigzag-labelled-binary-tree/
func pathInZigZagTree(label int) []int {
	level := int(math.Log2(float64(label)))
	elements := make([]int, level+1)
	elements[level] = label
	elements[0] = 1
	counter := level
	rest := label
	for counter > 1 {
		curr := rest - int(math.Pow(2, float64(counter)))
		parentIndex := curr / 2
		if counter%2 != 0 {
			rest = int(math.Pow(2, float64(counter))) - 1 - parentIndex
		} else {
			rest = 2*int(math.Pow(2, float64(counter-1))) - 1 - parentIndex
		}
		elements[counter-1] = rest
		counter--
	}

	return elements
}
