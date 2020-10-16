package algorythms

import (
	"math"
)

//https://leetcode.com/problems/group-the-people-given-the-group-size-they-belong-to/
func groupThePeople(groupSizes []int) [][]int {
	if len(groupSizes) == 0 {
		return make([][]int, 0)
	}
	mapping := make(map[int][]int, 0)
	min, max := math.MaxInt32, math.MinInt32
	for i, length := range groupSizes {
		min = int(math.Min(float64(min), float64(length)))
		max = int(math.Max(float64(max), float64(length)))

		if item, ok := mapping[length]; ok {
			mapping[length] = append(item, i)
		} else {
			newOne := []int{i}
			mapping[length] = newOne
		}
	}
	res := make([][]int, 0)
	for i:=min; i<=max; i++{
		if val, ok:=mapping[i]; ok{
			for j := 0; j < len(val)/i; j++ {
				res = append(res, val[i*j:i*j+i])
			}
		}
	}

	return res
}
