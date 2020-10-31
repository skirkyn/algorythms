package algorythms

import (
	"fmt"
	"testing"
)

func TestPrinting(t *testing.T) {
	fmt.Println(printTree(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}}))
}
