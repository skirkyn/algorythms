package algorythms

func reversePairs(nums []int) int {
	if nums == nil || len(nums) ==0 {
		return 0
	}
	counter := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i > j && nums[i]*2 < nums[j] {
				counter++
			}
		}
	}
	return counter
}
