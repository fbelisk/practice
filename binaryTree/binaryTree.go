package binaryTree


func FindAncestor(root *TreeNode, p *TreeNode, q *TreeNode) (*TreeNode, bool) {
	if root == nil {
		return nil, false
	}
	if root == p {
		return nil, true
	}
	if root == q {
		return nil, true
	}

	node1, findTarget1 := FindAncestor(root.Left, p, q)
	node2, findTarget2 :=  FindAncestor(root.Right, p, q)
	if node1 != nil {
		return node1, true
	}
	if node2 != nil {
		return node2, true
	}
	if findTarget1 && findTarget2 {
		return root, false
	}
	return nil, false
}

