package algorythms

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestVertical(t *testing.T) {
    res1:=verticalOrder(recoverTree(t, "3,9,20,null,null,15,7"))

    if "[[9] [3 15] [20] [7]]" != fmt.Sprint(res1){
    	t.Error("wrong result, expected [[9] [3 15] [20] [7]], got", res1)
	}
	res2:=verticalOrder(recoverTree(t, "3,9,8,4,0,1,7"))

	if "[[4] [9] [3 0 1] [8] [7]]" != fmt.Sprint(res2){
		t.Error("wrong result, expected [[4] [9] [3 0 1] [8] [7]], got", res2)
	}

	res3:=verticalOrder(recoverTree(t, "3,9,8,4,0,1,7,null,null,null,2,5"))

	if "[[4] [9 5] [3 0 1] [8 2] [7]]" != fmt.Sprint(res3){
		t.Error("wrong result, expected [[4] [9 5] [3 0 1] [8 2] [7]], got", res3)
	}
}

func recoverTree(t *testing.T, v string) *TreeNode{
	arr := strings.Split(v, ",")
	first, err:=strconv.Atoi(arr[0])
	if err != nil{
		t.Fatal(err)
	}
	root:=&TreeNode{Val: first}
	treeFromArray(arr, 0, root)
	return root
}
func treeFromArray(v []string, index int, node *TreeNode) {
	left := 2*index + 1
	right := 2*index + 2
	if left < len(v) {
		if v[left] != "null" {
			num, err := strconv.Atoi(v[left])
			if err == nil {
				node.Left = &TreeNode{Val: num}
				treeFromArray(v, left, node.Left)
			}
		}
	}
	if right < len(v) {
		if v[right] != "null" {
			num, err := strconv.Atoi(v[right])
			if err == nil {
				node.Right = &TreeNode{Val: num}
				treeFromArray(v, right, node.Right)
			}
		}
	}
}
