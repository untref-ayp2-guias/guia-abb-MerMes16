package ejercicios

import (
	"errors"
	"untref-ayp2/guia-abb/binarytree"
)

func PredecesorInOrder(bt *binarytree.BinarySearchTree[int], k int) (int, error) {
	nodo := bt.GetRoot()
	var predecesor *binarytree.BinaryNode[int]

	for nodo != nil {
		if k == nodo.GetData() {
			if nodo.GetLeft() != nil {
				max := max(nodo.GetLeft())
				return max.GetData(), nil
			}
			break
		} else if k > nodo.GetData() {
			predecesor = nodo
			nodo = nodo.GetRight()
		} else {
			nodo = nodo.GetLeft()
		}
	}

	if predecesor != nil {
		return predecesor.GetData(), nil
	}

	return 0, errors.New("no hay predecesores menores que el mínimo")
}

func max(n *binarytree.BinaryNode[int]) *binarytree.BinaryNode[int] {
	for n.GetRight() != nil {
		n = n.GetRight()
	}
	return n
}
