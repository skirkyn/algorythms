package algorythms

import (
	"testing"
)

func TestReverse(t *testing.T) {
	if reversePairs([]int{1,3,2,3,1}) != 2{
		t.Error("Number of pairs for 1,3,2,3,1should be 2")
	}
	if reversePairs([]int{2,4,3,5,1}) != 3{
		t.Error("Number of pairs for 1,3,2,3,1should be 3")
	}
}