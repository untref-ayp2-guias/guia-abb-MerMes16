package ejercicios

import (
	"errors"
	"untref-ayp2/guia-abb/binarytree"
	"untref-ayp2/guia-abb/stack"

	"golang.org/x/exp/constraints"
)

// BSTInOrderIterator es un iterador para recorrer un BinarySearchTree.
type BSTInOrderIterator[T constraints.Ordered] struct {
	internalStack *stack.Stack[*binarytree.BinaryNode[T]]
}

// NewBSTInOrderIterator crea un nuevo BSTInOrderIterator.
//
// Parámetros:
//   - `bst` un puntero a un BinarySearchTree.
//
// Retorna:
//   - un Iterator.
func NewBSTInOrderIterator[T constraints.Ordered](bst *binarytree.BinarySearchTree[T]) Iterator[T] {
	it := &BSTInOrderIterator[T]{
		internalStack: stack.NewStack[*binarytree.BinaryNode[T]](),
	}
	inOrder(it.internalStack, bst.GetRoot())

	return it
}

func inOrder[T constraints.Ordered](stack *stack.Stack[*binarytree.BinaryNode[T]], root *binarytree.BinaryNode[T]) *stack.Stack[*binarytree.BinaryNode[T]] {
	if root == nil {
		return stack
	}

	stack = inOrder(stack, root.GetRight())
	stack.Push(root)
	stack = inOrder(stack, root.GetLeft())

	return stack
}

// HasNext indica si hay un siguiente dato.
func (it *BSTInOrderIterator[T]) HasNext() bool {
	return !it.internalStack.IsEmpty()
}

// Next devuelve el siguiente dato, respetando el recorrido InOrder.
func (it *BSTInOrderIterator[T]) Next() (T, error) {
	if !it.HasNext() {
		var zero T
		err := errors.New("árbol vacío")
		return zero, err
	}
	node, _ := it.internalStack.Pop()

	return node.GetData(), nil
}
