package avl

import (
	"math/rand/v2"
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

func BenchmarkAvl_ContainsRandom(b *testing.B) {
	arr := make([]int, b.N)
	avl := New[int, int]()
	for i := 0; i < b.N; i++ {
		key := rand.Int()
		avl.Put(key, i)
		arr[i] = key
	}
	rand.Shuffle(b.N, func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		avl.Contains(arr[i])
	}
}

func BenchmarkAvl_ContainsSequence(b *testing.B) {
	arr := make([]int, b.N)
	avl := New[int, int]()
	for i := 0; i < b.N; i++ {
		avl.Put(i, i)
		arr[i] = i
	}
	rand.Shuffle(b.N, func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		avl.Contains(arr[i])
	}
}

func BenchmarkAvl_DeleteRandom(b *testing.B) {
	arr := make([]int, b.N)
	avl := New[int, int]()
	for i := 0; i < b.N; i++ {
		key := rand.Int()
		avl.Put(key, i)
		arr[i] = key
	}
	rand.Shuffle(b.N, func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		avl.Delete(arr[i])
	}
}

func BenchmarkAvl_DeleteSequence(b *testing.B) {
	arr := make([]int, b.N)
	avl := New[int, int]()
	for i := 0; i < b.N; i++ {
		avl.Put(i, i)
		arr[i] = i
	}
	rand.Shuffle(b.N, func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		avl.Delete(arr[i])
	}
}
