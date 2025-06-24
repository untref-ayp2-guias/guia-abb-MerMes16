package ejercicios

import (
	"errors"
	"untref-ayp2/guia-abb/binarytree"
	"untref-ayp2/guia-abb/stack"

	"golang.org/x/exp/constraints"
)

// BSTPostOrderIterator es un iterador para recorrer un BinarySearchTree en postorden.
type BSTPostOrderIterator[T constraints.Ordered] struct {
	internalStack *stack.Stack[*binarytree.BinaryNode[T]]
}

// NewBSTPostOrderIterator crea un nuevo BSTPostOrderIterator.
//
// Parámetros:
//   - `bst` un puntero a un BinarySearchTree.
//
// Retorna:
//   - un Iterator.
func NewBSTPostOrderIterator[T constraints.Ordered](bst *binarytree.BinarySearchTree[T]) Iterator[T] {
	it := &BSTPostOrderIterator[T]{
		internalStack: stack.NewStack[*binarytree.BinaryNode[T]](),
	}
	postOrder(it.internalStack, bst.GetRoot())

	return it
}

func postOrder[T constraints.Ordered](stack *stack.Stack[*binarytree.BinaryNode[T]], root *binarytree.BinaryNode[T]) *stack.Stack[*binarytree.BinaryNode[T]] {
	if root == nil {
		return stack
	}
	stack.Push(root)
	stack = postOrder(stack, root.GetRight())
	stack = postOrder(stack, root.GetLeft())

	return stack
}

// HasNext indica si hay un siguiente dato.
func (it *BSTPostOrderIterator[T]) HasNext() bool {
	return !it.internalStack.IsEmpty()
}

// Next devuelve el siguiente dato, respetando el recorrido PostOrder.
func (it *BSTPostOrderIterator[T]) Next() (T, error) {
	if !it.HasNext() {
		var zero T
		err := errors.New("árbol vacío")
		return zero, err
	}
	node, _ := it.internalStack.Pop()

	return node.GetData(), nil
}
