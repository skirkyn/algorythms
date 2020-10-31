package algorythms

import "testing"

func TestRecover(t *testing.T) {
	root1 := &TreeNode{Val: -1, Left: &TreeNode{Val: -1, Left: &TreeNode{Val: -1}, Right: &TreeNode{Val: -1}}, Right: &TreeNode{Val: -1}}
	findElements1 := Constructor(root1);
	if !findElements1.Find(1){
		t.Error("should find 1")
	}
	if !findElements1.Find(3){
		t.Error("should find 3")
	}
	if findElements1.Find(5){
		t.Error("should not find 5")
	}
	root2 := &TreeNode{Val: -1, Right: &TreeNode{Val: -1}}
	findElements2 := Constructor(root2);
	if findElements2.Find(1){
		t.Error("should not find 1")
	}
	if !findElements2.Find(2){
		t.Error("should find 2")
	}
}
