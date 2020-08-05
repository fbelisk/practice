package binaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var prev * TreeNode

func isValidBST(root *TreeNode) bool {
	return check(root)
}

//使用递归完成中序遍历
//先递归下探到最低层的左子节点，再找下一个节点，使用一个全局变量prev记录上一个节点，进行一次比较排序
func check(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !check(root.Left) {
		return false
	}

	if prev != nil && prev.Val >= root.Val {
		return false
	}
	prev = root
	if !check(root.Right) {
		return false
	}
	return true
}
