package ejercicios

import (
	"fmt"
	"untref-ayp2/guia-abb/binarytree"

	"golang.org/x/exp/constraints"
)

type TreeSet[T constraints.Ordered] struct {
	set *binarytree.BinarySearchTree[T]
}

func NewTreeSet[T constraints.Ordered](elements ...T) *TreeSet[T] {
	ts := &TreeSet[T]{set: binarytree.NewBinarySearchTree[T]()}
	for _, n := range elements {
		ts.Add(n)
	}
	return ts
}

func (ts *TreeSet[T]) Add(elements ...T) {
	for _, n := range elements {
		if !ts.set.Search(n) {
			ts.set.Insert(n)
		}
	}
}

func (ts *TreeSet[T]) Size() int {
	return ts.set.Size()
}

func (ts *TreeSet[T]) Contains(element T) bool {
	return ts.set.Search(element)
}

func (ts *TreeSet[T]) Remove(element T) {
	ts.set.Remove(element)
}

func (ts *TreeSet[T]) Values() []T {
	var arr []T
	return ts.recorrido(arr, ts.set.GetRoot())
}

func (ts *TreeSet[T]) recorrido(arr []T, n *binarytree.BinaryNode[T]) []T {
	if n == nil {
		return arr
	}

	arr = ts.recorrido(arr, n.GetLeft())
	arr = append(arr, n.GetData())
	arr = ts.recorrido(arr, n.GetRight())

	return arr
}

func (ts *TreeSet[T]) String() string {
	values := ts.Values()
	s := "Set: {"
	if ts.set.Size() != 0 {
		for _, n := range values {
			s += fmt.Sprintf("%v ", n)
		}
		s = s[0 : len(s)-1]
	}
	s += "}"
	return s
}
