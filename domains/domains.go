package domains

import (
	"fmt"
	"unsafe"
)

type SimpleStruct struct {
	X int
	Y int
	Z int
}

type TaggedStruct struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int64  `db:"score"`
}

type EmbeddedStruct struct {
	Active bool
}

type NumbersStruct struct {
	IntValue     int
	UintValue    uint
	Int64Value   int64
	Uint64Value  uint64
	Float64Value float64
	Int32Value   int32
	Uint32Value  uint32
	Float32Value float32
	Int16Value   int16
	Uint16Value  uint16
	Int8Value    int8
	Uint8Value   uint8
}

type MixedStruct struct {
	Flag   bool
	Name   string
	ID     int
	Active bool
}

type EmptyStruct struct {
}

type CommentedStruct struct {
	ID      string
	Value   int64
	Enabled bool // Some numeric value
	// Unique identifier

	// Status flag
}

// Struct with pointers
type PointerStruct struct {
	Name *string
	Age  *int
	Data *[]byte
}

func ShowMemory() {
	fmt.Printf("SimpleStruct size: %d bytes\n", unsafe.Sizeof(SimpleStruct{}))
	fmt.Printf("TaggedStruct size: %d bytes\n", unsafe.Sizeof(TaggedStruct{}))
	fmt.Printf("EmbeddedStruct size: %d bytes\n", unsafe.Sizeof(EmbeddedStruct{}))
	fmt.Printf("NumbersStruct size: %d bytes\n", unsafe.Sizeof(NumbersStruct{}))
	fmt.Printf("MixedStruct size: %d bytes\n", unsafe.Sizeof(MixedStruct{}))
	fmt.Printf("EmptyStruct size: %d bytes\n", unsafe.Sizeof(EmptyStruct{}))
	fmt.Printf("CommentedStruct size: %d bytes\n", unsafe.Sizeof(CommentedStruct{}))
	fmt.Printf("PointerStruct size: %d bytes\n", unsafe.Sizeof(PointerStruct{}))
}
