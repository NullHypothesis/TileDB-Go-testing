// This file was generated by sizegen. DO NOT EDIT.

package tiledb

import (
	"reflect"
	"unsafe"
)

// Constants for the byte size of each Go numeric type.
const (
	IntBytes        = uint64(unsafe.Sizeof(int(0)))
	UintBytes       = uint64(unsafe.Sizeof(uint(0)))
	UintptrBytes    = uint64(unsafe.Sizeof(uintptr(0)))
	BoolBytes       = uint64(unsafe.Sizeof(bool(false)))
	Int8Bytes       = uint64(unsafe.Sizeof(int8(0)))
	Int16Bytes      = uint64(unsafe.Sizeof(int16(0)))
	Int32Bytes      = uint64(unsafe.Sizeof(int32(0)))
	Int64Bytes      = uint64(unsafe.Sizeof(int64(0)))
	Uint8Bytes      = uint64(unsafe.Sizeof(uint8(0)))
	Uint16Bytes     = uint64(unsafe.Sizeof(uint16(0)))
	Uint32Bytes     = uint64(unsafe.Sizeof(uint32(0)))
	Uint64Bytes     = uint64(unsafe.Sizeof(uint64(0)))
	Float32Bytes    = uint64(unsafe.Sizeof(float32(0)))
	Float64Bytes    = uint64(unsafe.Sizeof(float64(0)))
	Complex64Bytes  = uint64(unsafe.Sizeof(complex64(0)))
	Complex128Bytes = uint64(unsafe.Sizeof(complex128(0)))
	ByteBytes       = Uint8Bytes
	RuneBytes       = Int32Bytes
)

// KindBytes maps each numeric reflect.Kind to the number of bytes it takes.
// Non-numeric kinds and aliases are not included.
var KindBytes = map[reflect.Kind]uint64{
	reflect.Int:        IntBytes,
	reflect.Uint:       UintBytes,
	reflect.Uintptr:    UintptrBytes,
	reflect.Bool:       BoolBytes,
	reflect.Int8:       Int8Bytes,
	reflect.Int16:      Int16Bytes,
	reflect.Int32:      Int32Bytes,
	reflect.Int64:      Int64Bytes,
	reflect.Uint8:      Uint8Bytes,
	reflect.Uint16:     Uint16Bytes,
	reflect.Uint32:     Uint32Bytes,
	reflect.Uint64:     Uint64Bytes,
	reflect.Float32:    Float32Bytes,
	reflect.Float64:    Float64Bytes,
	reflect.Complex64:  Complex64Bytes,
	reflect.Complex128: Complex128Bytes,
}
