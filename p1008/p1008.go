package main

import "fmt"

/* 1008. Construct Binary Search Tree from Preorder Traversal
 * Return the root node of a binary search tree
 * that matches the given preorder traversal.
 *
 * (Recall that a binary search tree is a binary tree where for every node,
 * any descendant of node.left has a value < node.val,
 * and any descendant of node.right has a value > node.val.
 *
 * Also recall that a preorder traversal displays the value of the node first,
 * then traverses node.left, then traverses node.right.)
 *
 * Example 1: Input: [8,5,1,7,10,12] Output: [8,5,10,1,7,null,12]
 *
 * Note:
 * 1 <= preorder.length <= 100
 * The values of preorder are distinct.
 */

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	head := TreeNode{Val: preorder[0]}

	if len(preorder) == 1 {
		return &head
	}

	if len(preorder) == 2 {
		if preorder[1] > head.Val {
			head.Right = bstFromPreorder([]int{preorder[1]})
		} else if preorder[1] < head.Val {
			head.Left = bstFromPreorder([]int{preorder[1]})
		}
		return &head
	}

	for i, v := range preorder[1:] {
		if v < head.Val {
			continue
		}
		head.Left = bstFromPreorder(preorder[1 : i+1])
		head.Right = bstFromPreorder(preorder[i+1:])
		return &head
	}

	head.Left = bstFromPreorder(preorder[1:])
	return &head
}

func (t *TreeNode) print() {
	if t == nil {
		return
	}
	fmt.Print(t.Val, " ")
	if t.Left != nil {
		t.Left.print()
	}
	if t.Right != nil {
		t.Right.print()
	}
}

func main() {
	fmt.Println("// '0' represent null value")

	a := []int{3, 1, 2}
	a_ans := []int{3, 1, 0, 0, 2}
	fmt.Println("Input : ", a)
	fmt.Println("Expect: ", a_ans)
	fmt.Print("Output:   ")
	b := bstFromPreorder(a)
	b.print()
}
