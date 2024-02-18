// Package avl provides a barebones key-value pair accepting Avl implementation and a
// Keyless variant for user-defined key generation based on provided values.
package avl

import (
	"cmp"
	"strings"
)

// Avl is a rigidly self-balancing bst.Bst. Every single operation in worst case is O(log n).
// Avl has fast lookups and generally more efficient operations thanks to the balanced nature
// of the tree. Unless inserting highly randomized data, this variant is recommended over regular
// bst.Bst. Implements obt.Obt, refer to Keyless for obt.Keyless variant.
type Avl[K cmp.Ordered, V any] struct {
	root *node[K, V]
	len  int
}

// New creates a new *Avl[K, V] with nil root and length equal to 0.
func New[K cmp.Ordered, V any]() *Avl[K, V] {
	return &Avl[K, V]{
		root: nil,
		len:  0,
	}
}

// Put implements obt.Obt.Put.
func (avl *Avl[K, V]) Put(key K, value V) (put bool) {
	if avl.root, put = avl.root.put(key, value); put {
		avl.len++
	}
	return put
}

// Delete implements obt.Obt.Delete.
func (avl *Avl[K, V]) Delete(key K) (deleted bool) {
	if avl.root, deleted = avl.root.delete(key); deleted {
		avl.len--
	}
	return deleted
}

// Contains implements obt.Obt.Contains.
func (avl *Avl[K, V]) Contains(key K) bool {
	cur := avl.root
	for cur != nil && key != cur.key {
		if key > cur.key {
			cur = cur.right
		} else {
			cur = cur.left
		}
	}
	return cur != nil && key == cur.key
}

// Len implements obt.Obt.Len.
func (avl *Avl[K, V]) Len() int {
	return avl.len
}

// String implements fmt.Stringer by showing the tree in a slice format,
// ignoring keys.
func (avl *Avl[K, V]) String() string {
	return "[" + strings.TrimSpace(avl.root.String()) + "]"
}

type (
	// KeyFunc is a key generating function signature for Avl variants with no explicit keys.
	KeyFunc[K cmp.Ordered, V any] func(V) K

	// Keyless is essentially a wrapper over a normal Avl with user-provided key generation
	// out of inserted values. The overhead is miniscule (~15-20 nanoseconds more per op
	// in comparison to the ordinary variant).
	Keyless[K cmp.Ordered, V any] struct {
		avl     *Avl[K, V]
		keyFunc KeyFunc[K, V]
	}
)

// NewKeyless constructs a *Keyless[Key, Value] with a key generating function for ordering.
func NewKeyless[K cmp.Ordered, V any](generateKey KeyFunc[K, V]) *Keyless[K, V] {
	return &Keyless[K, V]{
		avl:     New[K, V](),
		keyFunc: generateKey,
	}
}

// Put implements obt.KeylessObt.Put
func (k *Keyless[K, V]) Put(value V) (added bool) {
	return k.avl.Put(k.keyFunc(value), value)
}

// Delete implements obt.KeylessObt.Delete
func (k *Keyless[K, V]) Delete(value V) (removed bool) {
	return k.avl.Delete(k.keyFunc(value))
}

// Contains implements obt.KeylessObt.Contains
func (k *Keyless[K, V]) Contains(value V) bool {
	return k.avl.Contains(k.keyFunc(value))
}

// Len implements obt.KeylessObt.Len
func (k *Keyless[K, V]) Len() int {
	return k.avl.Len()
}

// String implements fmt.Stringer by showing the tree in a slice format,
// ignoring keys.
func (k *Keyless[K, V]) String() string {
	return k.avl.String()
}
