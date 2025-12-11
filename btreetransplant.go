package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return rplc
	}

	// If replacing the root node
	if root == node {
		if rplc != nil {
			rplc.Parent = nil
		}
		return rplc
	}

	// Find and update the parent's reference
	parent := node.Parent
	if parent != nil {
		if parent.Left == node {
			parent.Left = rplc
		} else if parent.Right == node {
			parent.Right = rplc
		}
	}

	// Update rplc's parent pointer
	if rplc != nil {
		rplc.Parent = parent
	}
	return root
}
