package lion_test

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/zzztttkkk/lion"
)

type unpreparedInt int32

type SetTest struct {
	A int
	B int64
	C unpreparedInt
}

func BenchmarkDirectlySet(b *testing.B) {
	var obj = SetTest{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		obj.A = 12
	}
}

func BenchmarkUnsafeUpdate(b *testing.B) {
	mptr := lion.Ptr[SetTest]()
	fieldOfA := lion.FieldOf[SetTest, lion.EmptyMeta](&mptr.A)

	var obj SetTest
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lion.UnsafeUpdate(&obj, fieldOfA, 12)
	}
}

func BenchmarkPtrSet(b *testing.B) {
	mptr := lion.Ptr[SetTest]()
	fieldOfA := lion.FieldOf[SetTest, lion.EmptyMeta](&mptr.A)

	var obj SetTest
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ptr := fieldOfA.PtrOfInstance(unsafe.Pointer(&obj)).(*int)
		*ptr = 12
	}
}

func BenchmarkChangeInstance(b *testing.B) {
	mptr := lion.Ptr[SetTest]()
	fieldOfA := lion.FieldOf[SetTest, lion.EmptyMeta](&mptr.A)

	var obj SetTest
	var objptr = unsafe.Pointer(&obj)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fieldOfA.ChangeInstance(objptr, 12)
	}
}

func BenchmarkChangeInstanceForUnpreparedType(b *testing.B) {
	mptr := lion.Ptr[SetTest]()
	fieldOfC := lion.FieldOf[SetTest, lion.EmptyMeta](&mptr.C)

	var obj SetTest
	var objptr = unsafe.Pointer(&obj)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fieldOfC.ChangeInstance(objptr, unpreparedInt(12))
	}
}

func BenchmarkNormalReflectSet(b *testing.B) {
	obj := &SetTest{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		objv := reflect.ValueOf(obj).Elem()
		fv := objv.Field(0)
		fv.Set(reflect.ValueOf(345))
	}
}
