package rbTree

type Color int8

const (
	RED   = Color(0)
	BLACK = Color(1)
)

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Color  Color
}

type RbTree struct {
	Root *Node
}

func (r *RbTree) Init() {
	node := &Node{nil, nil, nil, BLACK}
	r.Root = node
}

//
//          |                                  |
//          X                                  Y
//         / \         left rotate            / \
//        z  Y       ------------->         X   γ
//           / \                            / \
//          l  γ                         z  	l
//
//旋转操作都是以红色节点为中心
func (r *RbTree) LeftRotate(x *Node) {
	//左旋操作必须有右节点
	if x.Right == nil {
		return
	}

	y := x.Right
	p := x.Parent
	l := y.Left

	//中心节点变更
	x.Right = y.Left
	x.Parent = y

	//右节点变更
	y.Left = x
	y.Parent = p
	if l != nil {
		l.Parent = x
	}

	//变更节点的原父节点
	if p == nil {
		//根节点的话，根节点被右节点替换
		r.Root = y
	} else if x == p.Left {
		//当前节点是父节点的左子节点
		p.Left = y
	} else {
		//当前节点是父节点的右子节点
		p.Right = y
	}
}

//
//          |                                  |
//          X                                  Y
//         / \         right rotate           / \
//        Y   Z      ------------->          A   X
//       / \                                    / \
//      A  B                                 	B  Z
//
func (r *RbTree) RightRotate(x *Node) {
	if x.Left == nil {
		return
	}

	y := x.Left
	p := x.Parent
	b := y.Right

	//parent node modify
	if p == nil {
		r.Root = y
	} else if p.Left == x {
		p.Left = y
	} else {
		p.Right = y
	}

	//x node modify
	x.Left = y.Right
	x.Parent = y

	//y node modify
	y.Parent = p
	y.Right = x

	//b node modify
	b.Parent = x
}


