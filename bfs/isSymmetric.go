package bfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func isSymmetric(root *TreeNode) bool {
	return check(root.Left, root.Right)
}

func check(left, right *TreeNode) bool{
	if left == nil && right != nil {
		return false
	}
	if left != nil && right == nil {
		return false
	}
	if left == nil && right == nil {
		return true
	}
	return left.Val != right.Val && check(left.Left, right.Right) && check(left.Right, right.Left)
}
