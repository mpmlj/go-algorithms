package trees_and_graphs

import (
	"testing"
)

// 4.1
func TestRouteCheck(t *testing.T) {

	n1 := &Node{}
	_ = n1.Add(1)
	_ = n1.Add(2)
	n13 := n1.Add(3)
	_ = n1.Add(4)

	_ = n13.Add(11)
	_ = n13.Add(12)
	n13_3 := n13.Add(13)

	_ = n13_3.Add(21)
	_ = n13_3.Add(22)
	n3_3 := n13_3.Add(23)

	expected := true
	actual := n1.RouteCheck(n3_3)

	if expected != actual {
		t.Errorf("Expected: %v, got: %v", expected, actual)
	}
}

// 4.2 Minimal Tree
func TestMinimalTree(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	/**
	Expected tree:

				    5
			2		        7
		  1	      3	 	   6	     8
			     	 4		       9

	*/

	n := MinimalTree(arr)

	if n.Val != 5 {
		t.Errorf("Lvl 1. Expected %v, got %v ", 5, n.Val)
	}

	if n.Left.Val != 2 {
		t.Errorf("Lvl 2 - Left. Expected %v, got %v ", 2, n.Left.Val)
	}
	if n.Right.Val != 7 {
		t.Errorf("Lvl 2 - Right. Expected %v, got %v ", 7, n.Right.Val)
	}

	if n.Left.Left.Val != 1 {
		t.Errorf("Lvl 3.1 - Left. Expected %v, got %v ", 1, n.Left.Left.Val)
	}
	if n.Left.Right.Val != 3 {
		t.Errorf("Lvl 3.1 - Right. Expected %v, got %v ", 3, n.Left.Right.Val)
	}

	if n.Right.Left.Val != 6 {
		t.Errorf("Lvl 3.2 - Left. Expected %v, got %v ", 6, n.Right.Left.Val)
	}
	if n.Right.Right.Val != 8 {
		t.Errorf("Lvl 3.2 - Right. Expected %v, got %v ", 8, n.Right.Right.Val)
	}

	if n.Left.Right.Right.Val != 4 {
		t.Errorf("Lvl 4.1 - Left. Expected %v, got %v ", 4, n.Left.Right.Right.Val)
	}
	if n.Right.Right.Right.Val != 9 {
		t.Errorf("Lvl 4.1 - Right. Expected %v, got %v ", 9, n.Right.Right.Right.Val)
	}
}

// 4.3
func TestListOfDepths(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := MinimalTree(arr)

	listOfDepths := ListOfDepths(n)

	for i, l := range listOfDepths {
		t.Log("\n\n")
		t.Logf("Level %d\n", i+1)
		for n := l.head; n != nil; n = n.Next {
			t.Logf("%v -> ", n.Val)
		}
		t.Log("\n")
	}
}
