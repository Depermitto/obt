// Package bst provides a simple key-value pair accepting Bst implementation and a
// Keyless variant for user-defined key generation based on provided values.
package bst

import (
	"cmp"
	"fmt"
)

type (
	// Bst is the simplest implementation of a binary search tree. It provides
	// logarithmic O(log n) operation for insertion, deletion and search. This variant
	// provides no self-balancing, so worst-case (insertion of a sequence e.g. 1, 2, 3...)
	// the tree performs at the level of a linked list O(n).
	Bst[K cmp.Ordered, V any] struct {
		root *node[K, V]
		len  int
	}

	node[K cmp.Ordered, V any] struct {
		key   K
		value V
		left  *node[K, V]
		right *node[K, V]
	}
)

// New returns a new and empty *Bst[Key, Value].
func New[K cmp.Ordered, V any]() *Bst[K, V] {
	return &Bst[K, V]{
		root: nil,
		len:  0,
	}
}

func newLeaf[K cmp.Ordered, V any](key K, value V) *node[K, V] {
	return &node[K, V]{key: key, value: value}
}

func (bst *Bst[K, V]) search(key K) (found, parent *node[K, V]) {
	var (
		n *node[K, V] = bst.root
		p *node[K, V] = nil
	)
	for n != nil && key != n.key {
		p = n
		if key > n.key {
			n = n.right
		} else {
			n = n.left
		}
	}
	return n, p
}

// Put implements obt.Obt.Put.
func (bst *Bst[K, V]) Put(key K, value V) (added bool) {
	n, p := bst.search(key)
	switch {
	case n != nil:
		return false
	case p == nil:
		bst.root = newLeaf(key, value)
	case key > p.key:
		p.right = newLeaf(key, value)
	default:
		p.left = newLeaf(key, value)
	}
	bst.len++
	return true
}

func (bst *Bst[K, V]) resolve(parent, child, replacement *node[K, V]) {
	if parent == nil {
		bst.root = replacement
	} else if parent.left == child {
		parent.left = replacement
	} else {
		parent.right = replacement
	}
}

// Delete implements obt.Obt.Delete.
func (bst *Bst[K, V]) Delete(key K) (removed bool) {
	n, p := bst.search(key)
	switch {
	case n == nil:
		return false
	case n.left == nil && n.right == nil:
		bst.resolve(p, n, nil)
	case n.left == nil:
		bst.resolve(p, n, n.right)
	case n.right == nil:
		bst.resolve(p, n, n.left)
	default:
		var (
			t  *node[K, V] = n.right
			tp *node[K, V] = n
		)
		for t.left != nil {
			tp = t
			t = t.left
		}

		n.key, n.value = t.key, t.value
		bst.resolve(tp, t, nil)
	}
	bst.len--
	return true
}

// String implements fmt.Stringer by showing the tree in
// [(key1:value1)(key2:value2)...(keyN:valueN)] format.
func (bst *Bst[K, V]) String() (inOrder string) {
	if bst.root == nil {
		return "[]"
	}
	return fmt.Sprintf("[%v]", bst.root)
}

func (n *node[K, V]) String() (inOrder string) {
	if n.left != nil {
		inOrder += n.left.String()
	}
	inOrder += fmt.Sprintf("(%v:%v)", n.key, n.value)

	if n.right != nil {
		inOrder += n.right.String()
	}
	return inOrder
}

// Len implements obt.Obt.Len.
func (bst *Bst[K, V]) Len() int {
	return bst.len
}

// Contains implements obt.Obt.Contains.
func (bst *Bst[K, V]) Contains(key K) bool {
	n, _ := bst.search(key)
	return n != nil
}

type (
	// KeyFunc is a key generating function signature for Bst variants with no explicit keys.
	KeyFunc[K cmp.Ordered, V any] func(V) K

	// Keyless is essentially a wrapper over a normal Bst with user-provided key generation
	// out of inserted values. The overhead is miniscule (~15-20 nanoseconds more per op
	// in comparison to the ordinary Bst).
	Keyless[K cmp.Ordered, V any] struct {
		bst     *Bst[K, V]
		keyFunc KeyFunc[K, V]
	}
)

// NewKeyless constructs a *Keyless[Key, Value] with a key generating function out of values.
func NewKeyless[K cmp.Ordered, V any](generateKey KeyFunc[K, V]) *Keyless[K, V] {
	return &Keyless[K, V]{
		bst:     New[K, V](),
		keyFunc: generateKey,
	}
}

// Put implements obt.KeylessObt.Put
func (k *Keyless[K, V]) Put(value V) (added bool) {
	return k.bst.Put(k.keyFunc(value), value)
}

// Delete implements obt.KeylessObt.Delete
func (k *Keyless[K, V]) Delete(value V) (removed bool) {
	return k.bst.Delete(k.keyFunc(value))
}

// Contains implements obt.KeylessObt.Contains
func (k *Keyless[K, V]) Contains(value V) bool {
	return k.bst.Contains(k.keyFunc(value))
}

// Len implements obt.KeylessObt.Len
func (k *Keyless[K, V]) Len() int {
	return k.bst.Len()
}

// String implements fmt.Stringer by showing the tree in
// [(key1:value1)(key2:value2)...(keyN:valueN)] format.
func (k *Keyless[K, V]) String() string {
	return k.bst.String()
}
