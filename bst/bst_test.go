package bst

import (
	"fmt"
	"math/rand"
	"slices"
	"testing"
)

func TestNew(t *testing.T) {
	bst := New[int, int]()
	if bst == nil {
		t.Errorf("New[int]() returned nil")
	} else if bst.String() != "[]" {
		t.Errorf("expected %v; got %v", ``, bst.String())
	} else if bst.Len() != 0 {
		t.Errorf("expected %v; got %v", 0, bst.Len())
	}
}

func TestBst_Put(t *testing.T) {
	bst := New[int, int]()
	repr := "["
	for i := 0; i < 100; i++ {
		bst.Put(i, i)
		repr += fmt.Sprintf("(%v:%v)", i, i)
	}
	repr += "]"

	if bst.String() != repr {
		t.Errorf("expected %v; got %v", repr, bst.String())
	}
}

func TestBst_Remove(t *testing.T) {
	bst := New[int, int]()
	var nums []int
	for i := 0; i < 100; i++ {
		key := rand.Intn(1000)
		bst.Put(key, i)
		nums = append(nums, key)
	}
	slices.Sort(nums)
	nums = slices.Compact(nums)

	var removed []bool
	for _, num := range nums {
		removed = append(removed, bst.Delete(num))
	}

	if len(slices.Compact(removed)) != 1 {
		t.Errorf("didn't remove everything")
	}

	if bst.Len() != 0 {
		t.Errorf("expected %v; got %v", 0, bst.Len())
	}
}

func BenchmarkBst_PutRandom(b *testing.B) {
	arr := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		arr[i] = rand.Int()
	}

	b.ResetTimer()
	bst := New[int, int]()
	for i := 0; i < b.N; i++ {
		bst.Put(arr[i], i)
	}
}

func BenchmarkBst_PutSequence(b *testing.B) {
	bst := New[int, int]()
	for i := 0; i < b.N; i++ {
		bst.Put(i, i)
	}
}

func BenchmarkBst_ContainsRandom(b *testing.B) {
	arr := make([]int, b.N)
	bst := New[int, int]()
	for i := 0; i < b.N; i++ {
		key := rand.Int()
		bst.Put(key, i)
		arr[i] = key
	}
	rand.Shuffle(b.N, func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bst.Contains(arr[i])
	}
}

func BenchmarkBst_ContainsSequence(b *testing.B) {
	arr := make([]int, b.N)
	bst := New[int, int]()
	for i := 0; i < b.N; i++ {
		bst.Put(i, i)
		arr[i] = i
	}
	rand.Shuffle(b.N, func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bst.Contains(arr[i])
	}
}

func BenchmarkBst_DeleteRandom(b *testing.B) {
	arr := make([]int, b.N)
	bst := New[int, int]()
	for i := 0; i < b.N; i++ {
		key := rand.Int()
		bst.Put(key, i)
		arr[i] = key
	}
	rand.Shuffle(b.N, func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bst.Delete(arr[i])
	}
}

func BenchmarkBst_DeleteSequence(b *testing.B) {
	arr := make([]int, b.N)
	bst := New[int, int]()
	for i := 0; i < b.N; i++ {
		bst.Put(i, i)
		arr[i] = i
	}
	rand.Shuffle(b.N, func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bst.Delete(arr[i])
	}
}
