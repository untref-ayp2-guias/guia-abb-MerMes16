package ejercicios

import (
	"untref-ayp2/guia-abb/binarytree"
)

func IsBst(bt *binarytree.BinaryTree[int]) bool {
	if bt == nil {
		return true
	}
	return isBstByNode(bt.GetRoot())
}

func isBstByNode(n *binarytree.BinaryNode[int]) bool {
	if n == nil {
		return true
	}

	if (n.GetLeft() != nil && n.GetData() <= n.GetLeft().GetData()) || (n.GetRight() != nil && n.GetData() >= n.GetRight().GetData()) {
		return false
	}

	return isBstByNode(n.GetLeft()) && isBstByNode(n.GetRight())
}
