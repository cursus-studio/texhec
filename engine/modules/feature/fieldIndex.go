package feature

import (
	"reflect"
	"unsafe"
)

func GetFieldIndex[Type any, Field any](zero *Type, field *Field) int {
	// #nosec G103
	zeroPtr := uintptr(unsafe.Pointer(zero))
	// #nosec G103
	fieldPtr := uintptr(unsafe.Pointer(field))
	fieldOffset := fieldPtr - zeroPtr
	structType := reflect.TypeFor[Type]()
	for i := range structType.NumField() {
		field := structType.Field(i)
		if field.Offset == fieldOffset {
			return i
		}
	}
	panic("expected pointer to a field")
}
