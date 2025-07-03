package main

import (
	"testing"
	"unsafe"

	"github.com/orayew2002/art/art"
	"github.com/orayew2002/art/domains"
)

func TestStructSizes(t *testing.T) {
	t.Log("=== Sizes  ===")
	printSizes(t)

	total := getTotalSize()
	t.Logf("Total size: %d bytes", total)

	art.CleanStruct("domains/domains.go")

}

func printSizes(t testing.TB) {
	t.Logf("TaggedStruct size: %d bytes", unsafe.Sizeof(domains.TaggedStruct{}))
	t.Logf("MixedStruct size: %d bytes", unsafe.Sizeof(domains.MixedStruct{}))
	t.Logf("PointerStruct size: %d bytes", unsafe.Sizeof(domains.PointerStruct{}))
}

func getTotalSize() int {
	sizes := []uintptr{
		unsafe.Sizeof(domains.TaggedStruct{}),
		unsafe.Sizeof(domains.MixedStruct{}),
		unsafe.Sizeof(domains.PointerStruct{}),
	}

	var total int
	for _, size := range sizes {
		total += int(size)
	}
	return total
}
