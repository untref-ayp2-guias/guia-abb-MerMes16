package ejercicios

import (
	"testing"

	"untref-ayp2/guia-abb/binarytree"

	"github.com/stretchr/testify/assert"
)

/* Construct the following tree
          15
        /    \
       /      \
      10       20
     / \      /  \
    /   \    /    \
   8    12  16    25
*/

func TestPredecesorBst(t *testing.T) {
	bst := binarytree.NewBinarySearchTree[int]()
	bst.Insert(15)
	bst.Insert(10)
	bst.Insert(20)
	bst.Insert(8)
	bst.Insert(12)
	bst.Insert(16)
	bst.Insert(25)

	_, err := PredecesorInOrder(bst, 8)
	assert.EqualError(t, err, "no hay predecesores menores que el mínimo")
	p, _ := PredecesorInOrder(bst, 10)
	assert.Equal(t, 8, p)
	p, _ = PredecesorInOrder(bst, 12)
	assert.Equal(t, 10, p)
	p, _ = PredecesorInOrder(bst, 15)
	assert.Equal(t, 12, p)
	p, _ = PredecesorInOrder(bst, 16)
	assert.Equal(t, 15, p)
	p, _ = PredecesorInOrder(bst, 20)
	assert.Equal(t, 16, p)
	p, _ = PredecesorInOrder(bst, 25)
	assert.Equal(t, 20, p)
}

func TestNilPredecesorBst(t *testing.T) {
	bst := binarytree.NewBinarySearchTree[int]()

	_, err := PredecesorInOrder(bst, 16)
	assert.EqualError(t, err, "no hay predecesores menores que el mínimo")
}

func TestUnoSoloPredecesorBst(t *testing.T) {
	bst := binarytree.NewBinarySearchTree[int]()
	bst.Insert(100)

	_, err := PredecesorInOrder(bst, 16)
	assert.EqualError(t, err, "no hay predecesores menores que el mínimo")
}
