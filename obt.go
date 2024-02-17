// Package obt (Ordered Binary Tree) provides interfaces to implement by binary trees.
// It also provides convenience functions for many common operations.
package obt

import "cmp"

// Obt is short for Ordered Binary Tree. This essentially is an alias
// for Binary Search Tree, which is the simplest implementation  of this
// interface and some self-balancing Obts. All Obts are functionally Set-like.
//
// Obt and KeylessObt are designed to be mutually exclusive, you either implement
// one or the other.
type Obt[K cmp.Ordered, V any] interface {
	// Put inserts a key-value pair into the Obt.
	// Put must return true if the key-value pair has been added.
	Put(K, V) (put bool)

	// Delete removes a key-value pair from the Obt.
	// Delete must return true if the key-value pair has been deleted.
	Delete(K) (deleted bool)

	// Contains checks if the key exists in the Obt.
	Contains(K) bool

	// Len returns the amount of elements in the Obt.
	Len() int
}

// KeylessObt is a no-key version of Obt. The name may be misleading, as
// most implementations may still need to allocate for keys. The name KeylessObt
// comes from generating key-value pairs based on values (usually). See bst.Keyless
// for the simplest implementation.
type KeylessObt[V any] interface {
	// Put inserts a value (or a generated key-value pair) into the KeylessObt.
	// Put must return true if the value has been added.
	Put(V) (added bool)

	// Delete removes the value (or a generated key-value pair) from the KeylessObt.
	// Delete must return true if the value has been deleted.
	Delete(V) (removed bool)

	// Contains is like Obt.Contains but checks for values or generated keys.
	Contains(V) bool

	// Len returns the amount of elements in the KeylessObt.
	Len() int
}

// PutMany is a convenience function for putting multiple values into a KeylessObt.
// Ignores any errors that occured. The complexity is O(k log n) where k is the
// number of values to put, and n is the KeylessObt.Len of the tree.
func PutMany[V any](obt KeylessObt[V], value ...V) {
	for _, v := range value {
		obt.Put(v)
	}
}

// DeleteMany is like PutMany but for deletion.
func DeleteMany[V any](obt KeylessObt[V], value ...V) {
	for _, v := range value {
		obt.Delete(v)
	}
}

// ContainsAny is a convenience function for checking if any of multiple values
// is contained in KeylessObt. The complexity is O(k log n) where k is the number
// of values to check, and n is the KeylessObt.Len of the tree. This could technically
// be achieved in O(n) time, however for a couple of values, O(k log n) would be
// much faster.
//
// If you need to check every value iterate over the tree!
func ContainsAny[V any](obt KeylessObt[V], value ...V) bool {
	for _, v := range value {
		if obt.Contains(v) {
			return true
		}
	}
	return false
}

// ContainsAll is like ContainsAny, but checks if every value is contained in KeylessObt.
func ContainsAll[V any](obt KeylessObt[V], value ...V) bool {
	for _, v := range value {
		if !obt.Contains(v) {
			return false
		}
	}
	return true
}
