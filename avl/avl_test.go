package avl

import (
	"math/rand"
	"testing"
)

func BenchmarkAvl_PutRandom(b *testing.B) {
	arr := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		arr[i] = rand.Int()
	}

	b.ResetTimer()
	avl := New[int, int]()
	for i := 0; i < b.N; i++ {
		avl.Put(arr[i], i)
	}
}

func BenchmarkAvl_PutSequence(b *testing.B) {
	avl := New[int, int]()
	for i := 0; i < b.N; i++ {
		avl.Put(i, i)
	}
}

func BenchmarkAvl_Contains(b *testing.B) {
	avl := New[int, int]()
	for i := 0; i < b.N; i++ {
		avl.Put(rand.Int(), i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		avl.Contains(i)
	}
}
