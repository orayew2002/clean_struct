package main

import (
	"testing"
	"unsafe"

	"github.com/orayew2002/art/art"
	"github.com/orayew2002/art/domains"
)

func TestStructSizes(t *testing.T) {
	t.Log("=== Sizes BEFORE cleaning ===")
	beforeSizes := getSizes()
	printSizes(t, beforeSizes)

	art.CleanStruct("domains/domains.go")

	t.Log("=== Sizes AFTER cleaning ===")
	afterSizes := getSizes()
	printSizes(t, afterSizes)

	t.Log("=== Memory saved per struct ===")
	printSavedMemory(t, beforeSizes, afterSizes)
}

func BenchmarkCreateBefore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make([]domains.MixedStruct, 1_000_000)
	}
}

func BenchmarkCreateAfter(b *testing.B) {
	b.Skip("Replace with optimized struct when ready")
}

func getSizes() map[string]uintptr {
	return map[string]uintptr{
		"SimpleStruct":    unsafe.Sizeof(domains.SimpleStruct{}),
		"TaggedStruct":    unsafe.Sizeof(domains.TaggedStruct{}),
		"EmbeddedStruct":  unsafe.Sizeof(domains.EmbeddedStruct{}),
		"NumbersStruct":   unsafe.Sizeof(domains.NumbersStruct{}),
		"MixedStruct":     unsafe.Sizeof(domains.MixedStruct{}),
		"EmptyStruct":     unsafe.Sizeof(domains.EmptyStruct{}),
		"CommentedStruct": unsafe.Sizeof(domains.CommentedStruct{}),
		"PointerStruct":   unsafe.Sizeof(domains.PointerStruct{}),
	}
}

func printSizes(t testing.TB, sizes map[string]uintptr) {
	for name, size := range sizes {
		t.Logf("%-15s size: %d bytes", name, size)
	}
}

func printSavedMemory(t testing.TB, before, after map[string]uintptr) {
	var totalBefore, totalAfter uintptr
	for name := range before {
		b := before[name]
		a := after[name]
		saved := int64(b) - int64(a)
		t.Logf("%-15s saved: %d bytes", name, saved)
		totalBefore += b
		totalAfter += a
	}
	t.Logf("Total size BEFORE: %d bytes", totalBefore)
	t.Logf("Total size AFTER : %d bytes", totalAfter)
	t.Logf("TOTAL SAVED       : %d bytes", int64(totalBefore)-int64(totalAfter))
}
