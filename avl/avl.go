package avl

import (
	"cmp"
)

type Avl[K cmp.Ordered, V any] struct {
	root *node[K, V]
	len  int
}

func New[K cmp.Ordered, V any]() *Avl[K, V] {
	return &Avl[K, V]{
		root: nil,
		len:  0,
	}
}

func (avl *Avl[K, V]) Put(key K, value V) (put bool) {
	avl.root, put = avl.root.put(key, value)
	if put {
		avl.len++
	}
	return put
}

func (avl *Avl[K, V]) Delete(key K) (deleted bool) {
	//TODO implement me
	panic("implement me")
}

func (avl *Avl[K, V]) Contains(key K) bool {
	n := avl.root
	for n != nil && key != n.key {
		if key > n.key {
			n = n.right
		} else {
			n = n.left
		}
	}
	return n != nil && key == n.key
}

func (avl *Avl[K, V]) Len() int {
	return avl.root.height
}
