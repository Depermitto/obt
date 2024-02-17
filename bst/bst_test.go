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

func BenchmarkSlice_Add(b *testing.B) {
	var nums = make([]int, b.N)
	for i := 0; i < b.N; i++ {
		nums[i] = rand.Int()
	}
}

func BenchmarkBst_Add(b *testing.B) {
	bst := New[int, int]()
	for i := 0; i < b.N; i++ {
		bst.Put(rand.Int(), i)
	}
}

func BenchmarkKeyless_Add(b *testing.B) {
	bst := NewKeyless(func(i int) int { return i })
	for i := 0; i < b.N; i++ {
		bst.Put(rand.Int())
	}
}

func BenchmarkSlice_Remove(b *testing.B) {
	var nums = make([]int, b.N)
	for i := 0; i < b.N; i++ {
		nums[i] = rand.Int()
	}
	cont := slices.Clone(nums)
	rand.Shuffle(b.N, func(i, j int) {
		cont[i], cont[j] = cont[j], cont[i]
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		i := slices.Index(nums, cont[i])
		nums = append(nums[:i], nums[i+1:]...)
	}
}

func BenchmarkBst_Delete(b *testing.B) {
	var (
		bst  = New[int, int]()
		nums = make([]int, b.N)
	)
	for i := 0; i < b.N; i++ {
		key := rand.Int()
		bst.Put(key, i)
		nums[i] = key
	}
	rand.Shuffle(b.N, func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bst.Delete(nums[i])
	}
}

func BenchmarkKeyless_Delete(b *testing.B) {
	bst := NewKeyless(func(i int) int { return i })
	nums := make([]int, b.N)

	for i := 0; i < b.N; i++ {
		nums[i] = rand.Int()
		bst.Put(nums[i])
	}
	rand.Shuffle(b.N, func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bst.Delete(nums[i])
	}
}
