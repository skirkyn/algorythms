package algorythms

import (
	"fmt"
	"testing"
)

func TestGrouping(t *testing.T) {
	people1 := groupThePeople([]int{3, 3, 3, 3, 3, 1, 3})
	if "[[5] [0 1 2] [3 4 6]]" != fmt.Sprint(people1){
		t.Errorf("results are not equeal, expected: %s, actual: %s ", "[[5] [0 1 2] [3 4 6]]" , fmt.Sprint(people1))
	}
	people2 := groupThePeople([]int{2,1,3,3,3,2})
	if "[[1] [0 5] [2 3 4]]" != fmt.Sprint(people2){
		t.Errorf("results are not equeal, expected: %s, actual: %s ", "[[1] [0 5] [2 3 4]]" , fmt.Sprint(people2))
	}
	people3 := groupThePeople([]int{2,1,5,5,5,5,5,3,3,2,2,3,2,3,3,3})
	if "[[1] [0 9] [10 12] [7 8 11] [13 14 15] [2 3 4 5 6]]" != fmt.Sprint(people3) {
		t.Errorf("results are not equeal, expected: %s, actual: %s ", "[[1] [0 9] [10 12] [7 8 11] [13 14 15] [2 3 4 5 6]]", fmt.Sprint(people3))
	}
	people4 := groupThePeople([]int{1})
	if "[[0]]" != fmt.Sprint(people4) {
		t.Errorf("results are not equeal, expected: %s, actual: %s ", "[[0]]", fmt.Sprint(people4))
	}

	people5 := groupThePeople([]int{})
	if "[]" != fmt.Sprint(people5) {
		t.Errorf("results are not equeal, expected: %s, actual: %s ", "[[0]]", fmt.Sprint(people5))
	}

}
