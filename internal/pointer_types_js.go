package gd

import (
	"unsafe"

	"graphics.gd/internal/pointers"
)

type gdptr uint32

type Variant pointers.Trio[Variant]
type Signal pointers.Pair[Signal]
type Callable pointers.Pair[Callable]

type Dictionary pointers.Half[Dictionary]
type Array pointers.Half[Array]
type String pointers.Half[String]
type StringName pointers.Half[StringName]
type NodePath pointers.Half[NodePath]

type PackedByteArray pointers.Solo[PackedByteArray]
type PackedInt32Array pointers.Solo[PackedInt32Array]
type PackedInt64Array pointers.Solo[PackedInt64Array]
type PackedFloat32Array pointers.Solo[PackedFloat32Array]
type PackedFloat64Array pointers.Solo[PackedFloat64Array]
type PackedStringArray pointers.Solo[PackedStringArray]
type PackedVector2Array pointers.Solo[PackedVector2Array]
type PackedVector3Array pointers.Solo[PackedVector3Array]
type PackedVector4Array pointers.Solo[PackedVector4Array]
type PackedColorArray pointers.Solo[PackedColorArray]

type EnginePointer = uint32
type PackedPointers = [1]uint64

func UnsafeGet[T any](frame Address, index int) T {
	return *(*T)(Global.Memory.Index(frame, index, unsafe.Sizeof([1]T{})))
}

func UnsafeSet[T any](frame Address, value T) {
	ptr := Global.Memory.Index(frame, -1, unsafe.Sizeof([1]T{}))
	*(*T)(ptr) = value
	Global.Memory.Write(frame, ptr, unsafe.Sizeof([1]T{}))
}
