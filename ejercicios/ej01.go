package ejercicios

import (
	"errors"
	"untref-ayp2/guia-abb/binarytree"
)

func SecondLargestElement(bt *binarytree.BinarySearchTree[int]) (int, error) {
	if bt == nil {
		return 0, errors.New("No hay valores")
	}
	return findSecLargeElem(bt.GetRoot())
}

func findSecLargeElem(n *binarytree.BinaryNode[int]) (int, error) {
	if n == nil {
		return 0, errors.New("No hay valores")
	}

	right := n.GetRight()
	if left := n.GetLeft(); right == nil {
		if left == nil {
			return 0, errors.New("No hay valores")
		}
		return left.GetData(), nil
	}

	if right.GetRight() == nil {
		if rLeft := right.GetLeft(); rLeft != nil {
			return rLeft.GetData(), nil
		}
		return n.GetData(), nil
	}

	return findSecLargeElem(right)

}
