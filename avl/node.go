package avl

import (
	"cmp"
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
	return &node[K, V]{key: key, value: value, height: 1}
}

// getHeight is a safe height getter, on nil caller returns 0.
func (n *node[K, V]) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

// getBalance is a safe balance factor getter, on nil caller returns 0.
func (n *node[K, V]) getBalance() int {
	if n == nil {
		return 0
	}
	return n.right.getHeight() - n.left.getHeight()
}

func (n *node[K, V]) put(key K, value V) (*node[K, V], bool) {
	switch {
	case n == nil:
		return newLeaf(key, value), true
	case key < n.key:
		n.left, _ = n.left.put(key, value)
	case key > n.key:
		n.right, _ = n.right.put(key, value)
	default:
		return nil, false
	}
	n.height = 1 + max(n.left.getHeight(), n.right.getHeight())
	n.rebalance()
	return n, true
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
	n.left.height = 1 + max(n.left.left.getHeight(), n.left.right.getHeight())

	// new root balance
	n.height = 1 + max(n.left.getHeight(), n.right.getHeight())
}

func (n *node[K, V]) rightRot() {
	dangling := n.left.right
	old := *n

	*n = *n.left

	n.right = &old
	n.right.left = dangling

	// old root balance
	n.right.height = 1 + max(n.right.left.getHeight(), n.right.right.getHeight())

	// new root balance
	n.height = 1 + max(n.left.getHeight(), n.right.getHeight())
}
