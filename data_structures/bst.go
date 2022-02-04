package data_structures

type TreeNode struct {
	value int
	left	*TreeNode
	right *TreeNode
}

type BST struct {
	root	*TreeNode
}

func (tree *BST) Insert(value int) {
	newNode := &TreeNode {
		value: value,
		left: nil,
		right: nil,
	}

	if tree.root == nil {
		tree.root = newNode;
	} else {
		node := tree.root;
		for {
			if value >= node.value {
				if node.right == nil {
					node.right = newNode;
					break;
				} else {
					node = node.right;
				}
			} else {
				if node.left == nil {
					node.left = newNode;
					break;
				} else {
					node = node.left;
				}
			}
		}
	}
}

func (tree *BST) Lookup(value int) *TreeNode {
	node := tree.root;

	for node != nil && node.value != value {
		if value > node.value {
			node = node.right;
		} else {
			node = node.left;
		}
	}

	return node;
}

func (tree *BST) Remove(value int) {
	node := tree.root;
	var parent *TreeNode;
	// find the node while keeping track of the parent
	for node != nil {
		if value > node.value {
			parent = node;
			node = node.right;
		} else if value < node.value {
			parent = node;
			node = node.left;
		} else {
			if node.right == nil { // if no right child 
				if parent == nil {
					tree.root = node.left;
				} else {
					if value > parent.value { // deleting the right child of parent
						parent.right = node.left;
					} else { // deleting the left child of parent
						parent.left = node.left;
					}
				}
			} else if node.right.left == nil { // has right child but that right child has no left child
				if parent == nil {
					tree.root = node.left;
				} else {
					node.right.left = node.left;
					if value > parent.value {
						parent.right = node.right;
					} else {
						parent.left = node.left;
					}
				}
			} else { // has right child and that right child has left children
				leftMostChild := node.right.left;
				leftMostParent := node.right;

				for leftMostChild.left != nil {
					leftMostParent = leftMostChild;
					leftMostChild = leftMostChild.left;
				}
				
				leftMostChild.left = node.left;
				leftMostChild.right = node.right;
				leftMostParent.left = leftMostChild.right;
				if parent == nil {
					tree.root = node.left;
				} else {
					if value > parent.value {
						parent.right = leftMostChild;
					} else {
						parent.left = leftMostChild;
					}
				}				
			}
		}
	}
}