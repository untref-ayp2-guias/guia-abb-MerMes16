package ejercicios

import (
	"errors"
	"untref-ayp2/guia-abb/binarytree"
	"untref-ayp2/guia-abb/stack"

	"golang.org/x/exp/constraints"
)

// BSTPreOrderIterator es un iterador para recorrer un BinarySearchTree en preorden.
type BSTPreOrderIterator[T constraints.Ordered] struct {
	internalStack *stack.Stack[*binarytree.BinaryNode[T]]
}

// NewBSTPreOrderIterator crea un nuevo BSTPreOrderIterator.
//
// Parámetros:
//   - `bst` un puntero a un BinarySearchTree.
//
// Retorna:
//   - un Iterator.
func NewBSTPreOrderIterator[T constraints.Ordered](bst *binarytree.BinarySearchTree[T]) Iterator[T] {
	it := &BSTPreOrderIterator[T]{
		internalStack: stack.NewStack[*binarytree.BinaryNode[T]](),
	}
	preOrder(it.internalStack, bst.GetRoot())

	return it
}

func preOrder[T constraints.Ordered](stack *stack.Stack[*binarytree.BinaryNode[T]], root *binarytree.BinaryNode[T]) *stack.Stack[*binarytree.BinaryNode[T]] {
	if root == nil {
		return stack
	}

	stack = preOrder(stack, root.GetRight())
	stack = preOrder(stack, root.GetLeft())
	stack.Push(root)

	return stack

}

// HasNext indica si hay un siguiente dato.
func (it *BSTPreOrderIterator[T]) HasNext() bool {
	return !it.internalStack.IsEmpty()
}

// Next devuelve el siguiente dato, respetando el recorrido PreOrder.
func (it *BSTPreOrderIterator[T]) Next() (T, error) {
	if !it.HasNext() {
		var zero T
		err := errors.New("árbol vacío")
		return zero, err
	}
	node, _ := it.internalStack.Pop()

	return node.GetData(), nil
}
