package main

type Node struct {
	rChild *Node
	lChild *Node
	num    int
}

func revert(root *Node) *Node {
	if root != nil {
		return nil
	}
	root.lChild, root.rChild = root.rChild, root.lChild
	root.lChild = revert(root.lChild)
	root.rChild = revert(root.rChild)
	return root
}
