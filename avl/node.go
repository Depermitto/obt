package avl

import (
	"cmp"
	"fmt"
)

const (
	leftHeavy  int = -2
	rightHeavy int = 2
)

type node[K cmp.Ordered, V any] struct {
	key    K
	value  V
	left   *node[K, V]
	right  *node[K, V]
	height int
}

func newLeaf[K cmp.Ordered, V any](key K, value V) *node[K, V] {
	return &node[K, V]{
		key:    key,
		value:  value,
		height: 1,
	}
}

// getHeight is a safe height getter, on nil caller returns 0.
func (n *node[K, V]) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *node[K, V]) fixHeight() {
	n.height = 1 + max(n.left.getHeight(), n.right.getHeight())
}

// getBalance is a safe balance factor getter, on nil caller returns 0.
func (n *node[K, V]) getBalance() int {
	if n == nil {
		return 0
	}
	return n.right.getHeight() - n.left.getHeight()
}

func (n *node[K, V]) put(key K, value V) (newRoot *node[K, V], put bool) {
	switch {
	case n == nil:
		return newLeaf(key, value), true
	case key < n.key:
		n.left, put = n.left.put(key, value)
	case key > n.key:
		n.right, put = n.right.put(key, value)
	default:
		return nil, false
	}
	n.fixHeight()
	n.rebalance()
	return n, put
}

func (n *node[K, V]) delete(key K) (newRoot *node[K, V], deleted bool) {
	switch {
	case n == nil:
		return nil, false
	case key < n.key:
		n.left, deleted = n.left.delete(key)
	case key > n.key:
		n.right, deleted = n.right.delete(key)
	// key found
	case n.left == nil:
		return n.right, true
	case n.right == nil:
		return n.left, true
	default:
		var t *node[K, V] = n.left
		for t.right != nil {
			t = t.right
		}

		n.key, n.value = t.key, t.value
		n.left, deleted = n.left.delete(t.key)
	}
	n.fixHeight()
	n.rebalance()
	return n, deleted
}

func (n *node[K, V]) rebalance() {
	switch n.getBalance() {
	case leftHeavy:
		if n.left.getBalance() > 0 {
			n.left.leftRot()
		}
		n.rightRot()
	case rightHeavy:
		if n.right.getBalance() < 0 {
			n.right.rightRot()
		}
		n.leftRot()
	}
}

func (n *node[K, V]) leftRot() {
	dangling := n.right.left
	old := *n

	*n = *n.right

	n.left = &old
	n.left.right = dangling

	// old root balance
	n.left.fixHeight()
	// new root balance
	n.fixHeight()
}

func (n *node[K, V]) rightRot() {
	dangling := n.left.right
	old := *n

	*n = *n.left

	n.right = &old
	n.right.left = dangling

	// old root balance
	n.right.fixHeight()
	// new root balance
	n.fixHeight()
}

func (n *node[K, V]) String() (inOrder string) {
	if n.left != nil {
		inOrder += n.left.String()
	}
	inOrder += fmt.Sprintf("%v ", n.value)

	if n.right != nil {
		inOrder += n.right.String()
	}
	return inOrder
}
